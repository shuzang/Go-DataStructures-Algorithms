package sort

import (
	"fmt"
	"testing"
)

func ExampleInsertSort() {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(InsertSort(nums))
	//Output: [14 36 49 52 58 61 80]
}

func TestShellSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(ShellSort(nums))
}

func TestBubbleSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(BubbleSort(nums))
}

func TestSelectionSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(SelectionSort(nums))
}

func TestHeapSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(HeapSort(nums))
}

func TestQuickSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(QuickSort(nums, 0, len(nums)-1))
}

func TestMergeSort(t *testing.T) {
	nums := []int{52, 49, 80, 36, 14, 58, 61}
	fmt.Println(MergeSort(nums))
}
