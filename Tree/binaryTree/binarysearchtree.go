//Package binarysearchtree 实现了二叉搜索树数据结构
package binarysearchtree

import (
	"errors"
	"fmt"
)

//Item 树节点的数据可以是任意类型
type Item interface{}

//Node 二叉树节点结构
type Node struct {
	key   int   //以整型数据为例
	value Item  //节点数据可以是任意类型
	left  *Node //左儿子
	right *Node //右儿子
}

type basicFunction interface {
	Insert(key int, value Item)     //以构造二叉搜索树的方式插入节点
	PreOrderTraverse(f func(Item))  //前序遍历整棵树
	InOrderTraverse(f func(Item))   //中序遍历整棵树
	PostOrderTraverse(f func(Item)) //后序遍历整棵树
	Remove(key int) error           //移除树中等于给定值的节点
	Search(key int) (bool, error)   //如果树中存在与给定值相同的节点， 返回true,否则返回false
}

//BinarySearchTree 二叉树
type BinarySearchTree struct {
	size int //树中节点总数
	root *Node
}

//Insert 以构造二叉树的方式插入节点
func (bst *BinarySearchTree) Insert(key int, value Item) {
	newNode := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
	bst.size++
}

//寻找节点插入的合适位置
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

//PreOrderTraverse 前序遍历树中所有节点
func (bst *BinarySearchTree) PreOrderTraverse(f func(Item)) {
	preOrderTraverse(bst.root, f)
}

//前序遍历的递归函数
func preOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		f(node.value)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

//InOrderTraverse 中序遍历树中所有节点
func (bst *BinarySearchTree) InOrderTraverse(f func(Item)) {
	inOrderTraverse(bst.root, f)
}

//中序遍历的递归函数
func inOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}

//PostOrderTraverse 后序遍历树中所有节点
func (bst *BinarySearchTree) PostOrderTraverse(f func(Item)) {
	postOrderTraverse(bst.root, f)
}

//后序遍历的递归函数
func postOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.value)
	}
}

//Remove 移除树中与给定值相同的节点
func (bst *BinarySearchTree) Remove(key int) error {
	_, err := remove(bst.root, key)
	return err
}

//移除元素的递归函数
func remove(node *Node, key int) (*Node, error) {
	var err error
	switch {
	case node == nil:
		return nil, errors.New("error: empty tree")
	case key < node.key:
		node.left, err = remove(node.left, key)
		return node, err
	case key > node.key:
		node.right, err = remove(node.right, key)
		return node, err
	case node.left == nil && node.right == nil:
		node = nil
		return nil, nil
	case node.left == nil:
		node = node.right
		return node, nil
	case node.right == nil:
		node = node.left
		return node, nil
	}
	leftmostrightsubtree := node.right
	for {
		//寻找右子树中的最小值，位于右子树最左端
		if leftmostrightsubtree != nil && leftmostrightsubtree.left != nil {
			leftmostrightsubtree = leftmostrightsubtree.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightsubtree.key, leftmostrightsubtree.value
	node.right, err = remove(node.right, node.key)
	return node, err
}

//Search 搜索树中与给定值相同的节点，若存在，返回true
func (bst *BinarySearchTree) Search(key int) (bool, error) {
	return search(bst.root, key)
}

//搜索元素的递归函数
func search(node *Node, key int) (bool, error) {
	if node == nil {
		return false, errors.New("errors: item not exist")
	}
	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}
	return true, nil
}

// String prints a visual representation of the tree
func (bst *BinarySearchTree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}

//https://flaviocopes.com/golang-data-structure-binary-search-tree/
