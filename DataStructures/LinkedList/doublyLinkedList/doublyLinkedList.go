//Package doublylinkedlist is the implementation of doubly linked list using golang
//Note:the first node's "prev" pointer is nil, because its previous node is head node
//package doublylinkedlist

package main

import "fmt"

//Each node consist of at least three parts: value, prev and next pointer
type node struct {
	prev *node
	val  int
	next *node
}

//头节点，head指向第一个节点，size为节点数目
type doublyLinkedList struct {
	len  int
	next *node
}

//Initialize head and return the list address
func constructor() *doublyLinkedList {
	return &doublyLinkedList{0, nil}
}

//Insert node before the first node
func (ll *doublyLinkedList) headInsert(val int) {
	newNode := &node{nil, val, ll.next}
	if ll.len != 0 {
		ll.next.prev = newNode
	}
	ll.next = newNode
	ll.len++
}

//Insert node after the last node
func (ll *doublyLinkedList) tailInsert(val int) {
	newNode := &node{nil, val, nil}
	if ll.len == 0 {
		ll.next = newNode
		ll.len++
		return
	}
	cur := ll.next
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	newNode.prev = cur
	ll.len++
}

//Insert node before the index-th node
func (ll *doublyLinkedList) insert(val int, index int) {
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
	newNode := &node{cur, val, cur.next}
	cur.next.prev = newNode
	cur.next = newNode
	ll.len++
}

//Delete the index-th node
func (ll *doublyLinkedList) delete(index int) {
	if index < 1 || index > ll.len {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		ll.next = ll.next.next
		ll.next.prev = nil
		ll.len--
		return
	}

	cur := ll.next
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	if index == ll.len {
		cur.prev.next = nil
		cur.prev = nil
		ll.len--
		return
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	ll.len--
}

//Get the index-th node and return its value. if index invalid,return -1
func (ll *doublyLinkedList) get(index int) int {
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
func (ll *doublyLinkedList) printList() {
	cur := ll.next
	for ; cur.next != nil; cur = cur.next {
		fmt.Printf("%d<->", cur.val)
	}
	fmt.Println(cur.val)
}

func main() {

	sll := constructor()

	sll.tailInsert(2)
	sll.headInsert(11)
	sll.insert(3, 1)
	sll.delete(3)
	sll.printList()

	fmt.Println(sll.len)
	fmt.Println(sll.get(1))
}
