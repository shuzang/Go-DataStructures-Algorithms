//package linkedliststack use linkedlist to implement stack
//package linkedliststack
package main

import (
	"sync"
)

//node the type of the stack, actually it's a node of linkedlist
type node struct {
	value interface{}
	next  *node
}

//Stack the structure of stack
type Stack struct {
	length int
	lock   *sync.RWMutex
	next   *node
}

//NewStack creates a new Stack
func NewStack() *Stack {
	return &Stack{0, &sync.RWMutex{}, nil}
}

//IsEmpty return true if stack is empty,else false
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

//Push add an node to the top of the stack
func (s *Stack) Push(t interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.next = &node{t, s.next}
	s.length++
}

//Pop remove an node from the top of the stack
func (s *Stack) Pop() interface{} {
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

//Peek return top node of the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.next.value
}

/* func main() {
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for i := 0; i < 5; i++ {
		stack.Pop()
	}
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Peek())
} */
