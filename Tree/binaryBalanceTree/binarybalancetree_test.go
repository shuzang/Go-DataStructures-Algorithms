package binarybalancetree

import (
	"fmt"
	"testing"
)

func TestArrayToBinTree(t *testing.T) {
	arr := []Item{4, 2, 7, 1, 3}
	root := arrayToBinTree(arr)
	updateHeight(root)
	fmt.Println(levelorderTraversal(root))
}

func TestInsertIntoAVL(t *testing.T) {
	arr := []Item{4, 2, 7, 1, 3}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(insertIntoAVL(root, 5)))
}

func TestDeleteNode(t *testing.T) {
	arr := []Item{5, 3, 6, nil, nil, nil, 7}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(deleteNode(root, 3)))
}
