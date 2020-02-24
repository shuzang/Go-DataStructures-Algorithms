//Package binarysearchtree 实现了二叉搜索树数据结构
package binarybalancetree

import (
	"container/list"
	"reflect"
)

//Item：树节点的数据可以是任意类型
type Item interface{}

//TreeNode：结点结构
type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

type LevelOrder struct {
	Val    int
	Height int
}

//将传入数组转换成二叉树
func arrayToBinTree(nums []Item) *TreeNode {
	queue := list.New()
	if reflect.TypeOf(nums[0]) == nil {
		return nil
	}
	root := &TreeNode{nums[0].(int), 0, nil, nil}
	queue.PushBack(root)
	for i := 1; queue.Len() != 0; {
		root := queue.Remove(queue.Front()).(*TreeNode)
		if i < len(nums) {
			if reflect.TypeOf(nums[i]) != nil {
				root.Left = &TreeNode{nums[i].(int), 0, nil, nil}
				queue.PushBack(root.Left)
			}
			i++
		}
		if i < len(nums) {
			if reflect.TypeOf(nums[i]) != nil {
				root.Right = &TreeNode{nums[i].(int), 0, nil, nil}
				queue.PushBack(root.Right)
			}
			i++
		}
	}
	return root
}

func updateHeight(root *TreeNode) {
	update(root)
}

func update(root *TreeNode) int {
	var HL, HR int
	if root != nil {
		HL = update(root.Left)
		HR = update(root.Right)
		root.Height = max(HL, HR) + 1
		return max(HL, HR) + 1
	} else {
		return 0
	}
}

// 层次遍历
func levelorderTraversal(root *TreeNode) []LevelOrder {

	result := []LevelOrder{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		root = queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, LevelOrder{root.Val, root.Height})
		if root.Left != nil {
			queue.PushBack(root.Left)
		}
		if root.Right != nil {
			queue.PushBack(root.Right)
		}

	}
	return result
}

//搜索最小结点
func findMin(root *TreeNode) *TreeNode {
	if root != nil {
		for root.Left != nil {
			root = root.Left
		}
	}
	return root
}

//二叉树插入
func insertIntoAVL(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{val, 0, nil, nil}
	} else {
		if val > root.Val {
			root.Right = insertIntoAVL(root.Right, val)
		} else if val < root.Val {
			root.Left = insertIntoAVL(root.Left, val)
		}
	}
	updateHeight(root)
	return ajust(root)
}

//二叉树删除
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left != nil && root.Right != nil {
			root.Val = findMin(root.Right).Val
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			if root.Left == nil {
				root = root.Right
			} else if root.Right == nil {
				root = root.Left
			}
		}
	}
	updateHeight(root)
	return ajust(root)
}

func leftRotate(root *TreeNode) *TreeNode {
	tmp := root.Right
	root.Right = tmp.Left
	tmp.Left = root

	root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
	tmp.Height = max(getHeight(tmp.Left), getHeight(tmp.Right)) + 1
	return tmp
}

func rightThenLeftRotate(root *TreeNode) *TreeNode {
	tmp := rightRotate(root.Right)
	root.Right = tmp
	return leftRotate(root)
}

func rightRotate(root *TreeNode) *TreeNode {
	tmp := root.Left
	root.Left = tmp.Right
	tmp.Right = root

	root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
	tmp.Height = max(getHeight(tmp.Left), getHeight(tmp.Right)) + 1
	return tmp
}

func leftThenRightRotate(root *TreeNode) *TreeNode {
	tmp := leftRotate(root.Left)
	root.Left = tmp
	return rightRotate(root)
}

func ajust(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	compare := getHeight(root.Right) - getHeight(root.Left)
	if compare == 2 {
		if getHeight(root.Right.Right) > getHeight(root.Right.Left) {
			root = leftRotate(root)
		} else {
			root = rightThenLeftRotate(root)
		}
	} else if compare == -2 {
		if getHeight(root.Left.Left) > getHeight(root.Left.Right) {
			root = rightRotate(root)
		} else {
			root = leftThenRightRotate(root)
		}
	}
	return root
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Height
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
