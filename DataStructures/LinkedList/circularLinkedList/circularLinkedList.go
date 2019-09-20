/*circularlinkedlist can be singly linked list or doubly linked list.
  Now, We only implement two functions here:
   - Convert singly linked list into circular linked list
   - Circular linked list traversal
*/
//package circularlinkedlist
package main

import "fmt"

//Node consist of value and pointer to next
type node struct {
	val  int
	next *node
}

//Convert singly linked list into circular linked list
func convertCircular(singlylinkedlist *node) {
	cur := singlylinkedlist
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = singlylinkedlist
}

//Circular linked list traversal and print
func traversal(circularlinkedlist *node) {
	cur := circularlinkedlist
	for cur.next != circularlinkedlist {
		fmt.Printf("%v->", cur.val)
		cur = cur.next
	}
	fmt.Printf("%v->%v", cur.val, cur.next.val)
}

//Create a singly linked list for test
func createSingly(head *node) {
	cur := head
	for i := 2; i < 5; i++ {
		newNode := &node{i, nil}
		cur.next = newNode
		cur = cur.next
	}
}

func main() {
	ll := &node{1, nil}
	createSingly(ll)
	convertCircular(ll)
	traversal(ll)
}
