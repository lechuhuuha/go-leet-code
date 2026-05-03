package main

import (
	"os"
	"sync"
	"testing"
	"time"
)

type UnsafeUserSession struct {
	UserID     string
	LastSeenAt time.Time
	mu         sync.Mutex
}

func (*UnsafeUserSession) Touch(cache *sync.Map, userID string, now time.Time) bool {
	for {
		value, ok := cache.Load(userID)
		if !ok {
			return false
		}

		oldSession := value.(*UnsafeUserSession)
		newSession := &UnsafeUserSession{
			UserID:     oldSession.UserID,
			LastSeenAt: now,
		}

		if cache.CompareAndSwap(userID, oldSession, newSession) {
			return true
		}

	}
}

func (s *UnsafeUserSession) TouchUnsafe(now time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LastSeenAt = now
}

func (s *UnsafeUserSession) GetLastSeenAt() time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.LastSeenAt
}

// This test intentionally has a data race.
// Run it manually with:
//
// RUN_UNSAFE_RACE_TEST=1 go test -race -run TestUnsafeSyncMapPointerValueRace
//
// Expected result:
// WARNING: DATA RACE
func TestUnsafeSyncMapPointerValueRace(t *testing.T) {
	if os.Getenv("RUN_UNSAFE_RACE_TEST") != "1" {
		t.Skip("set RUN_UNSAFE_RACE_TEST=1 to run this intentional race test")
	}

	var cache sync.Map

	cache.Store("user-1", &UnsafeUserSession{
		UserID:     "user-1",
		LastSeenAt: time.Now(),
	})

	var wg sync.WaitGroup

	workers := 8
	loops := 10_000

	for workerID := 0; workerID < workers; workerID++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for i := 0; i < loops; i++ {
				value, ok := cache.Load("user-1")
				if !ok {
					t.Fatal("session not found")
				}

				session := value.(*UnsafeUserSession)

				session.Touch(&cache, "user-1", time.Unix(0, int64(workerID*loops+i)))
			}
		}(workerID)
	}

	wg.Wait()
}

// This test does NOT demonstrate a data race.
// The mutex protects LastSeenAt.
//
// It demonstrates a lifecycle bug:
// a goroutine can update an old *UserSession after that session
// was deleted/replaced in sync.Map.
func TestUnsafeSyncMapPointerValueCanUpdateDeletedSession(t *testing.T) {
	if os.Getenv("RUN_UNSAFE_LIFECYCLE_TEST") != "1" {
		t.Skip("set RUN_UNSAFE_LIFECYCLE_TEST=1 to run this lifecycle bug test")
	}

	var cache sync.Map

	oldInitialTime := time.Unix(100, 0)
	newInitialTime := time.Unix(200, 0)
	updatedAfterDeleteTime := time.Unix(999, 0)

	oldSession := &UnsafeUserSession{
		UserID:     "user-1",
		LastSeenAt: oldInitialTime,
	}

	cache.Store("user-1", oldSession)

	loadedSessionCh := make(chan *UnsafeUserSession, 1)
	allowUpdateCh := make(chan struct{})
	updateDoneCh := make(chan struct{})

	// Goroutine A:
	// Load the session first, but do not update it yet.
	go func() {
		value, ok := cache.Load("user-1")
		if !ok {
			t.Error("session not found")
			close(updateDoneCh)
			return
		}

		session := value.(*UnsafeUserSession)

		// Tell the test that we already loaded the pointer.
		loadedSessionCh <- session

		// Wait until another goroutine deletes/replaces the session.
		<-allowUpdateCh

		// This update is race-safe because Touch uses session.mu.
		// But it is logically wrong because this session was already deleted
		// from the cache.
		session.TouchUnsafe(updatedAfterDeleteTime)

		close(updateDoneCh)
	}()

	loadedSession := <-loadedSessionCh

	if loadedSession != oldSession {
		t.Fatal("expected goroutine to load the old session")
	}

	// Goroutine B / main test:
	// Delete old session from cache.
	deletedValue, loaded := cache.LoadAndDelete("user-1")
	if !loaded {
		t.Fatal("expected old session to be deleted")
	}

	deletedSession := deletedValue.(*UnsafeUserSession)
	if deletedSession != oldSession {
		t.Fatal("deleted session should be the old session")
	}

	// Create a new session for the same user.
	newSession := &UnsafeUserSession{
		UserID:     "user-1",
		LastSeenAt: newInitialTime,
	}

	cache.Store("user-1", newSession)

	// Now allow goroutine A to continue.
	// It still holds a pointer to oldSession.
	close(allowUpdateCh)

	<-updateDoneCh

	currentValue, ok := cache.Load("user-1")
	if !ok {
		t.Fatal("expected new session to exist in cache")
	}

	currentSession := currentValue.(*UnsafeUserSession)

	if currentSession != newSession {
		t.Fatal("cache should now point to the new session")
	}

	oldLastSeenAt := oldSession.GetLastSeenAt()
	newLastSeenAt := newSession.GetLastSeenAt()

	t.Logf("old deleted session LastSeenAt: %v", oldLastSeenAt)
	t.Logf("new current session LastSeenAt: %v", newLastSeenAt)

	if oldLastSeenAt != updatedAfterDeleteTime {
		t.Fatal("expected old deleted session to be updated after delete")
	}

	if newLastSeenAt == updatedAfterDeleteTime {
		t.Fatal("new session should not be updated by goroutine A")
	}
}
