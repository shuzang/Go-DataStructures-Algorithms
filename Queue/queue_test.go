package queue

import (
	"fmt"
	"testing"
)

var queue *Queue

func init() {
	queue = NewQueue()
}

func TestIsEmpty(t *testing.T) {
	fmt.Println(queue.IsEmpty())
}

func BenchmarkEnQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.EnQueue(i)
	}
}

func BenchmarkDnQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.EnQueue(i)
	}
	for i := 0; i < b.N; i++ {
		fmt.Println(queue.DnQueue())
	}
}

func BenchmarkPeek(b *testing.B) {
	queue.EnQueue("test")
	fmt.Println(queue.Peek())
}
