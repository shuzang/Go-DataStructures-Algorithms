package singlylinkedlist

import (
	"fmt"
	"testing"
)

//声明单链表
var linkedlist *singlyLinkedList

//所有测试函数执行前都需要初始化
func init() {
	linkedlist = &singlyLinkedList{0, nil}
}

func TestAddAtHead(t *testing.T) {
	linkedlist.AddAtHead(1)
	linkedlist.PrintList()
	// Output: 1
}

func TestAddAtTail(t *testing.T) {
	linkedlist.CreateList()
	linkedlist.PrintList()
	// Output: 0->1->2->3->4
	linkedlist.AddAtTail(3)
	linkedlist.PrintList()
	// Output: 0->1->2->3->4->3
}

func TestAddAtIndex(t *testing.T) {
	linkedlist.CreateList()
	linkedlist.PrintList()
	// Output: 0->1->2->3->4
	linkedlist.AddAtIndex(3, 3)
	linkedlist.PrintList()
	// Output: 0->1->3->2->3->4
}

func TestDeleteAtIndex(t *testing.T) {
	linkedlist.CreateList()
	linkedlist.PrintList()
	// Output: 0->1->2->3->4

	// //测试错误输出
	// if err := linkedlist.DeleteAtIndex(6); err != nil {
	// 	fmt.Printf("Error: %s\n", err)
	// }

	if err := linkedlist.DeleteAtIndex(5); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	linkedlist.PrintList()
	// Output: 0->1->2->3
}

func TestGet(t *testing.T) {
	linkedlist.CreateList()
	//item, err := linkedlist.Get(6) // 测试错误输出
	item, err := linkedlist.Get(3)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println(item)
	// Output: 2
}
