package heap

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	heap = append(heap, []int{20, 66, 43, 82, 72, 87}...)
	build(heap)
	fmt.Println(heap[1:])
}

func TestInsert(t *testing.T) {
	heap = append(heap, []int{56, 19, 40, 18, 9, 3}...)
	build(heap)
	fmt.Println(heap[1:])
	fmt.Println(Insert(heap, 72))
}

func TestDelete(t *testing.T) {
	heap = append(heap, []int{56, 19, 40, 18, 9, 3}...)
	build(heap)
	fmt.Println(heap[1:])
	fmt.Println(Delete(heap))
}
