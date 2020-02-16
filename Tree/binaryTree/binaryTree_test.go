package binarytree

import (
	"fmt"
	"testing"
)

// TestpreorderTraversal: 前序遍历测试
func TestPreorderTraversal(t *testing.T) {
	arr := []Item{1, nil, 2, 3}
	root := arrayToBinTree(arr)
	fmt.Println(preorderTraversal(root))
}

func TestInorderTraversal(t *testing.T) {
	arr := []Item{1, nil, 2, 3}
	root := arrayToBinTree(arr)
	fmt.Println(inorderTraversal(root))
}

func TestPostorderTraversal(t *testing.T) {
	arr := []Item{1, nil, 2, 3}
	root := arrayToBinTree(arr)
	fmt.Println(postorderTraversal(root))
}

func TestLevelorderTraversal(t *testing.T) {
	arr := []Item{3, 9, 20, nil, nil, 15, 7}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(root))
}
