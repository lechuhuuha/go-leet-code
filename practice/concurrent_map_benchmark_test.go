// concurrent_map_benchmark_test.go
package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

const keyCount = 1 << 16 // 65536

var sink int64

// 1. Read-heavy: sync.Map
// Workload: 99% reads, 1% writes.
// This simulates cache-like usage where most goroutines read,
// but some goroutines still refresh/update data.
func BenchmarkReadHeavySyncMap(b *testing.B) {
	var m sync.Map

	for i := 0; i < keyCount; i++ {
		m.Store(i, i)
	}

	var workerSeq int64

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		workerID := int(atomic.AddInt64(&workerSeq, 1))
		i := workerID * 9973

		var localSum int64

		for pb.Next() {
			readKey := i & (keyCount - 1)

			// 1% write
			if i%100 == 0 {
				// Write to a mostly disjoint key per worker.
				// This matches one common sync.Map-friendly pattern:
				// many readers, occasional independent writes.
				writeKey := keyCount + workerID
				m.Store(writeKey, i)
			} else {
				// 99% read
				value, ok := m.Load(readKey)
				if ok {
					// sync.Map trade-off:
					// value is any/interface{}, so we need type assertion.
					localSum += int64(value.(int))
				}
			}

			i++
		}

		atomic.AddInt64(&sink, localSum)
	})
}

// 2. Read-heavy: map + sync.RWMutex
// Same workload: 99% reads, 1% writes.
// Even though reads use RLock, every write still needs full Lock
// and can block new readers.
func BenchmarkReadHeavyRWMutexMap(b *testing.B) {
	m := make(map[int]int, keyCount)
	var mu sync.RWMutex

	for i := 0; i < keyCount; i++ {
		m[i] = i
	}

	var workerSeq int64

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		workerID := int(atomic.AddInt64(&workerSeq, 1))
		i := workerID * 9973

		var localSum int64

		for pb.Next() {
			readKey := i & (keyCount - 1)

			// 1% write
			if i%100 == 0 {
				writeKey := keyCount + workerID

				mu.Lock()
				m[writeKey] = i
				mu.Unlock()
			} else {
				// 99% read
				mu.RLock()
				value := m[readKey]
				mu.RUnlock()

				localSum += int64(value)
			}

			i++
		}

		atomic.AddInt64(&sink, localSum)
	})
}

// 3. Write-heavy: sync.Map
// Workload: 90% writes, 10% reads.
// This creates high mutation/churn, which is usually not ideal for sync.Map.
func BenchmarkWriteHeavySyncMap(b *testing.B) {
	var m sync.Map

	for i := 0; i < keyCount; i++ {
		m.Store(i, i)
	}

	var workerSeq int64

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		workerID := int(atomic.AddInt64(&workerSeq, 1))
		i := workerID * 9973

		var localSum int64

		for pb.Next() {
			key := i & (keyCount - 1)

			// 10% read
			if i%10 == 0 {
				value, ok := m.Load(key)
				if ok {
					localSum += int64(value.(int))
				}
			} else {
				// 90% write/delete churn
				if i%2 == 0 {
					m.Store(key, i)
				} else {
					m.Delete(key)
				}
			}

			i++
		}

		atomic.AddInt64(&sink, localSum)
	})
}

// 4. Write-heavy: map + sync.RWMutex
// Same workload: 90% writes, 10% reads.
// For this kind of frequent mutation, normal map + explicit locking
// is often simpler and can be faster.
func BenchmarkWriteHeavyRWMutexMap(b *testing.B) {
	m := make(map[int]int, keyCount)
	var mu sync.RWMutex

	for i := 0; i < keyCount; i++ {
		m[i] = i
	}

	var workerSeq int64

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		workerID := int(atomic.AddInt64(&workerSeq, 1))
		i := workerID * 9973

		var localSum int64

		for pb.Next() {
			key := i & (keyCount - 1)

			// 10% read
			if i%10 == 0 {
				mu.RLock()
				value := m[key]
				mu.RUnlock()

				localSum += int64(value)
			} else {
				// 90% write/delete churn
				mu.Lock()
				if i%2 == 0 {
					m[key] = i
				} else {
					delete(m, key)
				}
				mu.Unlock()
			}

			i++
		}

		atomic.AddInt64(&sink, localSum)
	})
}

// Case 1: all goroutines update the same key.
// This removes the "disjoint keys" advantage of sync.Map.
func BenchmarkHotKeyWriteSyncMap(b *testing.B) {
	var m sync.Map
	m.Store(0, 0)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Store(0, i)
			i++
		}
	})
}

func BenchmarkHotKeyWriteRWMutexMap(b *testing.B) {
	m := map[int]int{0: 0}
	var mu sync.RWMutex

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.Lock()
			m[0] = i
			mu.Unlock()
			i++
		}
	})
}

// Case 2: small keyspace.
// Many goroutines fight over only 16 keys.
func BenchmarkSmallKeyspaceWriteSyncMap(b *testing.B) {
	var m sync.Map

	for i := 0; i < 16; i++ {
		m.Store(i, i)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i & 15
			m.Store(key, i)
			i++
		}
	})
}

func BenchmarkSmallKeyspaceWriteRWMutexMap(b *testing.B) {
	m := make(map[int]int, 16)
	var mu sync.RWMutex

	for i := 0; i < 16; i++ {
		m[i] = i
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i & 15

			mu.Lock()
			m[key] = i
			mu.Unlock()

			i++
		}
	})
}

// Case 3: single goroutine write-heavy.
// This checks raw overhead without concurrency benefit.
func BenchmarkSingleGoroutineWriteSyncMap(b *testing.B) {
	var m sync.Map
	m.Store(0, 0)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := i & 15
		m.Store(key, i)
	}
}

func BenchmarkSingleGoroutineWriteRWMutexMap(b *testing.B) {
	m := make(map[int]int, 16)
	var mu sync.RWMutex

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := i & 15

		mu.Lock()
		m[key] = i
		mu.Unlock()
	}
}
