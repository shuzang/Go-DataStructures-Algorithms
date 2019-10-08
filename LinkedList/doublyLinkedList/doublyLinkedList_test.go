package doublylinkedlist

import (
	"fmt"
	"testing"
)

var linkedlist *doublyLinkedList

func init() {
	linkedlist = &doublyLinkedList{0, nil}
}

func BenchmarkHeadInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linkedlist.headInsert(i)
	}
}

func BenchmarkTailInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linkedlist.tailInsert(i)
	}
}

func TestInsert(t *testing.T) {
	for i := 0; i < 10; i++ {
		linkedlist.tailInsert(i)
	}
	linkedlist.insert("test", 3)
}

func TestDelete(t *testing.T) {
	for i := 10; i < 20; i++ {
		linkedlist.tailInsert(i)
	}
	linkedlist.delete(5)
}

func TestGet(t *testing.T) {
	for i := 20; i < 30; i++ {
		linkedlist.tailInsert(i)
	}
	fmt.Println(linkedlist.get(7))
}

func TestPrintList(t *testing.T) {
	for i := 30; i < 40; i++ {
		linkedlist.tailInsert(i)
	}
	linkedlist.printList()
}
