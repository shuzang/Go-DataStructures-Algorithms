package linkedstack

import (
	"sync"
)

//node the type of the LinkedStack, actually it's a node of linkedlist
type node struct {
	value interface{}
	next  *node
}

//LinkedStack the structure of LinkedStack
type LinkedStack struct {
	length int
	lock   *sync.RWMutex
	next   *node
}

//NewLinkedStack creates a new LinkedStack
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{0, &sync.RWMutex{}, nil}
}

//IsEmpty return true if LinkedStack is empty,else false
func (s *LinkedStack) IsEmpty() bool {
	return s.length == 0
}

//Push add an node to the top of the LinkedStack
func (s *LinkedStack) Push(t interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.next = &node{t, s.next}
	s.length++
}

//Pop remove an node from the top of the LinkedStack
func (s *LinkedStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	tmp := s.next
	s.next = s.next.next
	s.length--
	return tmp.value
}

//Peek return top node of the LinkedStack
func (s *LinkedStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.next.value
}
