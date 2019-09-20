package circularlinkedlist

import "testing"

var linkedlist *node

func init() {
	linkedlist = &node{"test", nil}
	cur := linkedlist
	for i := 2; i < 10; i++ {
		newNode := &node{i, nil}
		cur.next = newNode
		cur = cur.next
	}
}

func TestConvertCircular(t *testing.T) {
	convertCircular(linkedlist)
}

func TestTraversal(t *testing.T) {
	traversal(linkedlist)
}
