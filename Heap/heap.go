package heap

var heap = make([]int, 1)

func Insert(heap []int, key int) []int {
	heap = append(heap, key)
	i := len(heap) - 1
	for i > 1 && heap[i] > heap[i/2] {
		heap[i], heap[i/2] = heap[i/2], heap[i]
		i /= 2
	}
	return heap[1:]
}

func Delete(heap []int) (int, []int) {
	var parent, child int
	Max, t := heap[1], len(heap)-1
	heap[1], heap[t] = heap[t], heap[1]
	heap = heap[:t]
	for parent = 1; parent*2 < len(heap); parent = child {
		child = parent * 2
		if child+1 < len(heap) && heap[child+1] > heap[child] {
			child++
		}
		if heap[child] <= heap[parent] {
			break
		}
		heap[child], heap[parent] = heap[parent], heap[child]
	}
	return Max, heap[1:]
}

//假设传入的数组已经按顺序填好元素
func build(heap []int) {
	var parent, child int
	for i := (len(heap) - 1) / 2; i >= 1; i-- {
		for parent = i; parent*2 < len(heap); parent = child {
			child = parent * 2
			if child+1 < len(heap) && heap[child+1] > heap[child] {
				child++
			}
			if heap[child] <= heap[parent] {
				break
			}
			heap[child], heap[parent] = heap[parent], heap[child]
		}
	}
}
