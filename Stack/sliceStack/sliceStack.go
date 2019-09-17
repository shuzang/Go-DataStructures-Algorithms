//package slicestack use slice to implement stack
//package slicestack
package main

import (
	"fmt"
	"sync"
)

//Item the type of the stack
type Item interface{}

//ItemStack the structure of stack
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

//NewStack creates a new ItemStack
func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []Item{}
	return s
}

//IsEmpty return true if stack is empty,else false
func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}

//Push add an Item to the top of the stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

//Pop remove an Item from the top of the stack
func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

//Peek return top Item of the stack
func (s *ItemStack) Peek() Item {
	return s.items[len(s.items)-1]
}

func main() {
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for i := 0; i < 5; i++ {
		stack.Pop()
	}
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Peek())
}
