//Package arraystack : use array implement stack, but in golang ,actually use slice
package arraystack

import (
	"sync"
)

//Item the type of the ArrayStack
type Item interface{}

//ArrayStack the structure of ArrayStack
type ArrayStack struct {
	items []Item
	lock  sync.RWMutex
}

//NewStack creates a new ArrayStack
func NewStack() *ArrayStack {
	s := &ArrayStack{}
	s.items = []Item{}
	return s
}

//IsEmpty return true if ArrayStack is empty,else false
func (s *ArrayStack) IsEmpty() bool {
	return len(s.items) == 0
}

//Push add an Item to the top of the ArrayStack
func (s *ArrayStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

//Pop remove an Item from the top of the ArrayStack
func (s *ArrayStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

//Peek return top Item of the ArrayStack
func (s *ArrayStack) Peek() Item {
	return s.items[len(s.items)-1]
}
