package binarysearchtree

import "testing"

var bst BinarySearchTree

func fillTree(bst *BinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func TestInsert(t *testing.T) {
	fillTree(&bst)
	bst.String()
	bst.Insert(11, "11")
	bst.String()
}
