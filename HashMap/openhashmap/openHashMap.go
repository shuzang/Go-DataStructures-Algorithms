/*Package openhashmap 是开地址哈希的实现，由于golang中map结构本身在底层
就是使用哈希表实现的，所以我们并不使用map结构来实现这一数据结构，
冲突解决的办法使用双散列，因为它的产生的元素分布比其他几种办法如
线性探测或平方探测都好。*/
package openhashmap

import "fmt"

type openHashMap struct {
	bucketsize int //槽位数量
	itemsize   int //元素数目
	bucket     []int
}

//Hash 使用双散列解决冲突
func (h *openHashMap) Hash(item int, i int) int {
	var h1, h2 int
	h1 = item % h.bucketsize
	if item%2 == 0 {
		h2 = item%(h.bucketsize-1) + 1
	} else {
		h2 = item%(h.bucketsize-2) + 1
	}
	val := (h1 + i*h2) % h.bucketsize
	return val
}

//Init 初始化哈希表
func (h *openHashMap) Init(bucketsize int) {
	h.bucket = []int{}
	h.bucketsize = bucketsize
	for i := 0; i < h.bucketsize; i++ {
		h.bucket = append(h.bucket, -1)
	}
	h.itemsize = 0
}

//Destory 销毁哈希表
func (h *openHashMap) Destory() {
	h.bucket = []int{}
	h.bucketsize, h.itemsize = 0, 0
}

//Insert 插入元素
func (h *openHashMap) Insert(item int) {
	var i int
	if h.itemsize == h.bucketsize {
		fmt.Println("bucket has no space!")
	}
	// 相同的键不允许重复插入
	if h.Lookup(item) {
		fmt.Println("item already exist!")
	}
	for ; i < h.bucketsize; i++ {
		if h.bucket[h.Hash(item, i)] == -1 {
			h.bucket[h.Hash(item, i)] = item
			h.itemsize++
		}
	}
	if i == h.bucketsize {
		fmt.Println("Hash functiom error!")
	}
}

//Remove 删除给定元素
func (h *openHashMap) Remove(item int) {
	var i int
	for ; i < h.bucketsize; i++ {
		if h.bucket[h.Hash(item, i)] == -1 {
			fmt.Println("item not exist!")
		} else if h.bucket[h.Hash(item, i)] == item {
			h.bucket[h.Hash(item, i)] = -1
			h.itemsize--
		}
	}
	if i == h.bucketsize {
		fmt.Println("Hash function error!")
	}
}

//Lookup 查找与给定值匹配的元素
func (h *openHashMap) Lookup(item int) bool {
	var i int
	for ; i < h.bucketsize; i++ {
		if h.bucket[h.Hash(item, i)] == -1 {
			return false
		} else if h.bucket[h.Hash(item, i)] == item {
			return true
		}
	}
	return false
}
