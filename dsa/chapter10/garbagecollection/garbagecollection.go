package garbagecollection

import "sync"

type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

func NewReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}

type Stack struct {
	references []*ReferenceCounter
	Count      int
}

func (s *Stack) New() {
	s.references = make([]*ReferenceCounter, 0)
}

func (s *Stack) Push(r *ReferenceCounter) {
	s.references = append(s.references[:s.Count], r)
	s.Count++
}

func (s *Stack) Pop() *ReferenceCounter {
	if s.Count == 0 {
		return nil
	}

	l := len(s.references)
	references := s.references[l-1]
	if l > 1 {
		s.references = s.references[:l-1]
	} else {
		s.references = s.references[0:]
	}

	s.Count = len(s.references)
	return references
}
