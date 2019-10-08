/*Package circularlinkedlist can be singly linked list or doubly linked list.
  Now, We only implement two functions here:
   - Convert singly linked list into circular linked list
   - Circular linked list traversal*/
package circularlinkedlist

import "fmt"

//Node consist of value and pointer to next
type node struct {
	val  interface{}
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
	fmt.Printf("%v->%v\n", cur.val, cur.next.val)
}
