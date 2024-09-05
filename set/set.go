package set

import (
	"fmt"
	"iter"
)

// Set contains unique elements
// Elements must be comparable
type Set[E comparable] struct {
	m map[E]struct{}
}

// New creates an empty Set
func New[E comparable]() *Set[E] {
	return &Set[E]{
		m: make(map[E]struct{}),
	}
}

// Add inserts a new element into set, if element exists it is not duplicated
func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// Contains checks if a given element is in the set
func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}

// Push iterates over elements in the set and applies the given function to each
func (s *Set[E]) Push(f func(E) bool) {
	for v := range s.m {
		if !f(v) {
			return
		}
	}
}

// Pull returns two functions: first returns a next element in the set when called, second stops iteration and cleans up
func (s *Set[E]) Pull() (func() (E, bool), func()) {
	ch := make(chan E)
	stopCh := make(chan bool)

	go func() {
		defer close(ch)
		for v := range s.m {
			select {
			case ch <- v:
			case <-stopCh:
				return
			}
		}
	}()

	next := func() (E, bool) {
		v, ok := <-ch
		return v, ok
	}

	stop := func() {
		close(stopCh)
	}

	return next, stop
}

// All returns an iterator over elements in the set
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

// Union creates a new Set from the given sets
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	r := New[E]()
	for v := range s1.m {
		r.Add(v)
	}
	for v := range s2.m {
		r.Add(v)
	}
	return r
}

// PrintAllElementsPush prints all elements in the set using push iterator
func PrintAllElementsPush[E comparable](s *Set[E]) {
	s.Push(func(v E) bool {
		fmt.Print(v, " ")
		return true
	})
	fmt.Println()
}

// PrintAllElementsPull prints all elements in the set using pull iterator
func PrintAllElementsPull[E comparable](s *Set[E]) {
	next, stop := s.Pull()
	defer stop()
	for v, ok := next(); ok; v, ok = next() {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
