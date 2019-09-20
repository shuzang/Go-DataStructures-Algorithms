package arraystack

import (
	"fmt"
	"testing"
)

var stack *ArrayStack

func init() {
	stack = NewStack()
}

func BenchmarkPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		fmt.Println(stack.Pop())
	}
}

func TestPeek(t *testing.T) {
	stack.Push("It's a test")
	fmt.Println(stack.Peek())
}
