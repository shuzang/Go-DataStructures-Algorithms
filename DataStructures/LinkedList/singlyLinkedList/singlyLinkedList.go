//Package singlylinkedlist is the implementation of singly linked list using golang
//package singlylinkedlist

package main

import "fmt"

//Each node consists of at least two parts: value and the pointer to the next node
type node struct {
	val  int
	next *node
}

//head node
type singlyLinkedList struct {
	len  int
	next *node
}

//Initialize the head and return the list address
func constructor() *singlyLinkedList {
	return &singlyLinkedList{0, nil}
}

//Insert node before the first node
func (ll *singlyLinkedList) headInsert(val int) {
	tmp := &node{val, ll.next}
	ll.next = tmp
	ll.len++
}

//Insert node after the last node
func (ll *singlyLinkedList) tailInsert(val int) {
	tmp := &node{val, nil}

	if ll.next == nil {
		ll.next = tmp
		ll.len++
		return
	}

	cur := ll.next
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = tmp
	ll.len++
}

//Insert node before the index-th node
func (ll *singlyLinkedList) insert(val int, index int) {
	if index < 1 || index > ll.len {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		ll.headInsert(val)
		return
	}
	cur := ll.next
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	tmp := &node{val, cur.next}
	cur.next = tmp
	ll.len++
}

//Delete the index-th node
func (ll *singlyLinkedList) delete(index int) {
	if index < 1 || index > ll.len {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		ll.next = ll.next.next
		ll.len--
		return
	}

	cur := ll.next
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	ll.len--
}

//Get the index-th node and return its value. if index invalid,return -1
func (ll *singlyLinkedList) get(index int) int {
	if index < 1 || index > ll.len {
		fmt.Println("index invalid!")
		return -1
	}
	cur := ll.next
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

//Print the list
func (ll *singlyLinkedList) printList() {
	cur := ll.next
	for ; cur.next != nil; cur = cur.next {
		fmt.Printf("%d->", cur.val)
	}
	fmt.Println(cur.val)
}

func main() {

	sll := constructor()

	sll.tailInsert(2)
	sll.headInsert(11)
	sll.insert(3, 2)
	sll.delete(1)
	sll.printList()

	fmt.Println(sll.len)
	fmt.Println(sll.get(1))
}
