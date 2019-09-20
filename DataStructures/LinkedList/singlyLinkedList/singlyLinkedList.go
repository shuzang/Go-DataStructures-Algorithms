package singlylinkedlist

import "fmt"

//node consists of two parts: value and the pointer to the next node
type node struct {
	val  interface{}
	next *node
}

//head node
type singlyLinkedList struct {
	length int
	head   *node
}

//Initialize the head and return the list address
func constructor() *singlyLinkedList {
	return &singlyLinkedList{0, nil}
}

//Insert node before the first node
func (list *singlyLinkedList) headInsert(val interface{}) {
	tmp := &node{val, list.head}
	list.head = tmp
	list.length++
}

//Insert node after the last node
func (list *singlyLinkedList) tailInsert(val interface{}) {
	tmp := &node{val, nil}

	if list.head == nil {
		list.head = tmp
		list.length++
		return
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = tmp
	list.length++
}

//Insert node before the index-th node
func (list *singlyLinkedList) insert(val interface{}, index int) {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		list.headInsert(val)
		return
	}
	cur := list.head
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	tmp := &node{val, cur.next}
	cur.next = tmp
	list.length++
}

//Delete the index-th node
func (list *singlyLinkedList) delete(index int) {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		list.head = list.head.next
		list.length--
		return
	}

	cur := list.head
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	list.length--
}

//Get the index-th node and return its value. if index invalid,return -1
func (list *singlyLinkedList) get(index int) interface{} {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return -1
	}
	cur := list.head
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

//Print the list
func (list *singlyLinkedList) printList() {
	cur := list.head
	for ; cur.next != nil; cur = cur.next {
		fmt.Printf("%v->", cur.val)
	}
	fmt.Println(cur.val)
}
