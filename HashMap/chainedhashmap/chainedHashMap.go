package chainedhashmap

import "fmt"

type node struct {
	val  interface{}
	next *node
}

type linkedlist struct {
	len  int
	next *node
}

type chainedHashMap struct {
	cap    int
	bucket []*linkedlist
}

//Init 初始化链式哈希表
func (h *chainedHashMap) Init(cap int) {
	h.cap = cap
	if h.cap != 0 {
		h.bucket = make([]*linkedlist, h.cap, h.cap)
	}
	for _, v := range h.bucket {
		v = new(linkedlist)
		v.len, v.next = 0, nil
	}
}

//Destroy 销毁链式哈希表
func (h *chainedHashMap) Destroy() {
	for _, v := range h.bucket {
		v.len, v.next = 0, nil
	}
	h.bucket = nil
}

//Hash 暂时假设传入的元素为数字，使用除留取余法构造哈希函数
func (h *chainedHashMap) Hash(item int) int {
	return item % h.cap
}

//Lookup 查找元素是否已存在
func (h *chainedHashMap) Lookup(item int) bool {
	key := h.Hash(item)
	if h.bucket[key].len == 0 {
		return false
	}
	tmp := h.bucket[key].next
	for ; tmp != nil; tmp = tmp.next {
		if tmp.val == item {
			return true
		}
	}
	return false
}

//Insert 向表中插入元素
func (h *chainedHashMap) Insert(item int) {
	if h.Lookup(item) {
		fmt.Println("item already exist!")
	}
	key := h.Hash(item)
	tmp := h.bucket[key].next
	if tmp == nil {
		h.bucket[key].next = &node{item, nil}
		h.bucket[key].len++
		return
	}
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &node{item, nil}
	h.bucket[key].len++
}

//Remove 删除表中指定元素
func (h *chainedHashMap) Remove(item int) {
	if !h.Lookup(item) {
		fmt.Println("Can't remove because item not exist!")
	}
	key := h.Hash(item)
	tmp := h.bucket[key].next
	if tmp.val == item {
		h.bucket[key].next = tmp.next
	}
	for tmp.next.val != item {
		tmp = tmp.next
	}
	tmp.next = tmp.next.next
	h.bucket[key].len--
}

//Size 返回哈希表元素总个数
func (h *chainedHashMap) Size() int {
	var sum int
	for _, v := range h.bucket {
		sum += v.len
	}
	return sum
}
