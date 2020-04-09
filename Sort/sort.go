package sort

/*@description: 插入排序，最重要的一部分是数组元素后移
  @author: 2020-03-31
  @param: nums []int 待排序数组
  #@return: _ []int 已排序数组
*/
func InsertSort(nums []int) []int {
	var i, j, tmp int // tmp用来临时存储本轮待比较元素
	for i = 1; i < len(nums); i++ {
		tmp = nums[i] // 临时存储本轮待比较元素
		for j = i; j > 0 && tmp < nums[j-1]; j-- {
			nums[j] = nums[j-1] // 数组元素后移
		}
		nums[j] = tmp // 在正确的位置插入元素
	}
	return nums
}

func ShellSort(nums []int) []int {
	var i, j int
	// dk为增量，初始为待排序数组长度的一半，然后每轮循环减半
	for dk := len(nums) / 2; dk > 0; dk /= 2 {
		// 内层循环是间隔为dk的插入排序
		for i = dk; i < len(nums); i++ {
			tmp := nums[i]
			for j = i; j >= dk && tmp < nums[j-dk]; j -= dk {
				nums[j] = nums[j-dk]
			}
			nums[j] = tmp
		}
	}
	return nums
}

/* @description: 冒泡排序，依次比较相邻的两个数
   @author: shuzang 2020-03-31
   @param: nums []int 待排序数组
   @return: _ []int 已排序数组
*/
func BubbleSort(nums []int) []int {
	var flag bool
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag { //如果当前数组已有序，跳出循环
			break
		}
	}
	return nums
}

/* @description: 冒泡排序，依次比较相邻的两个数
   @author: shuzang 2020-03-31
   @param: nums []int 待排序数组
   @return: _ []int 已排序数组
*/
func SelectionSort(nums []int) []int {
	var min int
	for i := 0; i < len(nums)-1; i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}

/* @description: 堆排序，堆的建立与调整使用标准库函数
   @author: shuzang 2020-04-02
   @param: nums []int 待排序数组
   @return: _ []int 已排序数组
*/
func HeapSort(nums []int) []int {
	result := []int{}
	// 初始化堆
	for i := (len(nums) - 1) / 2; i >= 1; i-- {
		nums = heapAdjust(nums, i)
	}
	//每次取出根元素并重新调整堆
	for len(nums)-1 != 0 {
		t := len(nums) - 1
		nums[1], nums[t] = nums[t], nums[1]
		result = append([]int{nums[t]}, result...)
		nums = nums[:t]
		nums = heapAdjust(nums, 1)
	}
	return result
}

/* @description: 堆调整函数，调整数组使其符合堆的要求
   @author: shuzang 2020-04-04
   @param: nums []int 传入的待调整数组，start int 调整的元素下标
   @return: _ []int 调整好的数组

*/
func heapAdjust(nums []int, start int) []int {
	var parent, child int
	for parent = start; parent*2 < len(nums); parent = child {
		child = parent * 2
		if child+1 < len(nums) && nums[child+1] > nums[child] {
			child++
		}
		if nums[child] <= nums[parent] {
			break
		}
		nums[child], nums[parent] = nums[parent], nums[child]
	}
	return nums
}

/* @description: 快速排序的递归解法
   @author: shuzang 2020-04-04
   @param: nums []int 待排序数组，low int 第一个指向标记，high int 第二个指向标记
		   初始参数传入时，low和high应该为第一个元素和最后一个元素
   @return: _ []int 完成排序或部分完成排序的数组的数组
*/
func QuickSort(nums []int, low, high int) []int {
	if low < high {
		pivot := nums[low]
		i, j := low, high
		for i < j {
			for ; i < j && nums[j] >= pivot; j-- {
			}
			nums[i], nums[j] = nums[j], nums[i]
			for ; i < j && nums[i] <= pivot; i++ {
			}
			nums[j], nums[i] = nums[i], nums[j]
		}
		nums[i] = pivot
		QuickSort(nums, low, i-1)
		QuickSort(nums, i+1, high)
	}
	return nums
}

/*@description: 归并两个子序列
  @author: shuzang 2020-04-04
  @param: left, right []int 待合并的两个子序列，
  @return: _ []int 合并后的子序列
*/
func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	for m < len(left) && n < len(right) {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	result = append(append(result, left[m:]...), right[n:]...)
	return result
}

/*@description: 归并排序的递归解法
  @author: shuzang 2020-04-04
  @param: nums []int 待排序数组
  @return: _ []int 已排序数组
*/
/* func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := MergeSort(nums[:i])
	right := MergeSort(nums[i:])
	result := merge(left, right)
	return result
} */

/*@description: 归并排序的非递归解法
  @author: shuzang 2020-04-04
  @param: nums []int 待排序数组
  @return: _ []int 已排序数组
*/
func MergeSort(nums []int) []int {
	for i := 1; i <= len(nums); i *= 2 {
		var j int
		for ; j <= len(nums)-2*i; j += 2 * i {
			result := merge(nums[j:i+j], nums[i+j:i*2+j])
			copy(nums[j:i*2+j], result)
		}
		if i+j < len(nums) {
			copy(nums[j:], merge(nums[j:i+j], nums[i+j:]))
		}
	}
	return nums
}
