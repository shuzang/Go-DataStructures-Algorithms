package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestSearchBST(t *testing.T) {
	arr := []Item{4, 2, 7, 1, 3}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(searchBST(root, 2)))
}

func TestInsertIntoBST(t *testing.T) {
	arr := []Item{4, 2, 7, 1, 3}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(insertIntoBST(root, 5)))
}

func TestDeleteNode(t *testing.T) {
	arr := []Item{5, 3, 6, 2, 4, nil, 7}
	root := arrayToBinTree(arr)
	fmt.Println(levelorderTraversal(deleteNode(root, 3)))
}
