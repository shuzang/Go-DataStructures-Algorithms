//Package doublylinkedlist is the implementation of doubly linked list using golang
//Note:the first node's "prev" pointer is nil, because its previous node is head node
package doublylinkedlist

import "fmt"

//node consist of three parts: value, prev and next pointer
type node struct {
	prev *node
	val  interface{}
	next *node
}

//头节点，head指向第一个节点，size为节点数目
type doublyLinkedList struct {
	length int
	next   *node
}

//Initialize head and return the list address
func constructor() *doublyLinkedList {
	return &doublyLinkedList{0, nil}
}

//Insert node before the first node
func (list *doublyLinkedList) headInsert(val interface{}) {
	newNode := &node{nil, val, list.next}
	if list.length != 0 {
		list.next.prev = newNode
	}
	list.next = newNode
	list.length++
}

//Insert node after the last node
func (list *doublyLinkedList) tailInsert(val interface{}) {
	newNode := &node{nil, val, nil}
	if list.length == 0 {
		list.next = newNode
		list.length++
		return
	}
	cur := list.next
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	newNode.prev = cur
	list.length++
}

//Insert node before the index-th node
func (list *doublyLinkedList) insert(val interface{}, index int) {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		list.headInsert(val)
		return
	}
	cur := list.next
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	newNode := &node{cur, val, cur.next}
	cur.next.prev = newNode
	cur.next = newNode
	list.length++
}

//Delete the index-th node
func (list *doublyLinkedList) delete(index int) {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return
	}
	if index == 1 {
		list.next = list.next.next
		list.next.prev = nil
		list.length--
		return
	}

	cur := list.next
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	if index == list.length {
		cur.prev.next = nil
		cur.prev = nil
		list.length--
		return
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	list.length--
}

//Get the index-th node and return its value. if index invalid,return -1
func (list *doublyLinkedList) get(index int) interface{} {
	if index < 1 || index > list.length {
		fmt.Println("index invalid!")
		return -1
	}
	cur := list.next
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

//Print the list
func (list *doublyLinkedList) printList() {
	cur := list.next
	for ; cur.next != nil; cur = cur.next {
		fmt.Printf("%v<->", cur.val)
	}
	fmt.Println(cur.val)
}
