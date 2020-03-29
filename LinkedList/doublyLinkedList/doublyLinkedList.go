package doublylinkedlist

import (
	"errors"
	"fmt"
)

//双链表结点
type node struct {
	prev *node // 指向上一个结点的指针
	val  int   //结点的值
	next *node // 指向下一个结点的指针
}

//头节点，也是双链表的起始
type doublyLinkedList struct {
	length int   // 链表长度
	next   *node // 指向第一个结点的指针
}

/* @description: 初始化链表(头结点)
   @author: shuzang 2020-03-29
   @param: 无
   @return: _ *doublyLinkedList 指向双链表(头结点)的指针
*/
func constructor() *doublyLinkedList {
	return &doublyLinkedList{0, nil}
}

/* @description: 在链表的第一个元素之前插入结点，插入后，新结点将成为链表的第一个结点
   @author: shuzang	2020-03-29
   @param: val	int	要添加的元素的值
   @return: 无
*/
func (list *doublyLinkedList) AddAtHead(val int) {
	newNode := &node{nil, val, list.next}
	if list.length != 0 {
		list.next.prev = newNode
	}
	list.next = newNode
	list.length++
}

/* @description: 将新结点追加到链表的最后一个元素
   @author: shuzang 2020-03-29
   @param: val int 要添加的元素的值
   @return: 无
*/
func (list *doublyLinkedList) AddAtTail(val int) {
	newNode := &node{nil, val, nil}
	if list.length == 0 {
		list.next = newNode
	} else {
		cur := list.next
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
		newNode.prev = cur
	}

	list.length++
}

/* @description: 在链表的第 index 个结点前添加结点，插入后新结点成为第 index 个结点，
如果 index 大于链表长度，结点添加到链表末尾，如果 index 小于 1，则在头部插入结点
   @author: shuzang 2020-03-29
   @param: index int 要插入的位置，起始数字为 1; val int 要插入的元素的值
   @return: 无
*/
func (list *doublyLinkedList) AddAtIndex(index, val int) {
	if index > list.length {
		list.AddAtTail(val)
	} else if index <= 1 {
		list.AddAtHead(val)
	} else {
		cur := list.next
		for i := 1; i < index-1; i++ {
			cur = cur.next
		}
		newNode := &node{cur, val, cur.next}
		cur.next.prev = newNode
		cur.next = newNode
	}

	list.length++
}

/* @description: 如果索引有效，删除链表第 index 个位置的结点
   @author: shuzang 2020-03-29
   @param: index int 要删除的元素位置
   @return: _ error 索引无效时返回错误
*/
func (list *doublyLinkedList) DeleteAtIndex(index int) error {
	if index < 1 || index > list.length {
		return errors.New("param - index is invalid")
	} else if index == 1 {
		list.next = list.next.next
		if list.next != nil {
			list.next.prev = nil
		}

	} else {
		cur := list.next
		for i := 1; i < index-1; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		if cur.next != nil {
			cur.next.prev = cur
		}
	}

	list.length--
	return nil
}

/* @description: 如果索引有效，获取链表中第 index 个结点的值
   @author: shuzang 2020-03-29
   @param: index int 要获取的元素位置
   @return: _ int 获取元素的值; _ error 索引无效时返回错误
*/
func (list *doublyLinkedList) Get(index int) (int, error) {
	if index < 1 || index > list.length {
		return -1, errors.New("param - index is invalid")
	}
	cur := list.next
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur.val, nil
}

/* @description: 工具函数，打印单链表
   @author: shzuang 2020-03-29
   @param: 无
   @return: 无
*/
func (list *doublyLinkedList) PrintList() {
	cur := list.next
	fmt.Println("当前双链表为: ")
	for ; cur.next != nil; cur = cur.next {
		fmt.Printf("%v<->", cur.val)
	}
	fmt.Println(cur.val)
}

/* @description: 工具函数，建立一条简单的单链表用作删除和插入测试
   @author: shzuang 2020-03-28
   @param: 无
   @return: 无
*/
func (list *doublyLinkedList) CreateList() {
	for i := 0; i < 5; i++ {
		list.AddAtTail(i)
	}
}
