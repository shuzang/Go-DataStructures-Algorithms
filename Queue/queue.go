package queue

import "sync"

//Qnode to store queue item
type Qnode struct {
	val  interface{}
	next *Qnode
}

//Queue structure of the queue
type Queue struct {
	length      int
	lock        *sync.RWMutex
	front, rear *Qnode
}

//NewQueue create a new queue
func NewQueue() *Queue {
	return &Queue{0, &sync.RWMutex{}, nil, nil}
}

//IsEmpty return true if Queue is empty,else false
func (s *Queue) IsEmpty() bool {
	return s.length == 0
}

//EnQueue add an item to the queue
func (s *Queue) EnQueue(t interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	tmp := &Qnode{t, nil}
	if s.IsEmpty() {
		s.front = tmp
		s.rear = tmp
		s.length++
		return
	}
	s.rear.next = tmp
	s.rear = tmp
	s.length++
}

//DnQueue remove an item from the queue
func (s *Queue) DnQueue() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	tmp := s.front
	s.front = s.front.next
	if s.front == nil {
		s.rear = nil
	}
	s.length--
	return tmp.val
}

//Peek return the first item of queue
func (s *Queue) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.front.val
}
