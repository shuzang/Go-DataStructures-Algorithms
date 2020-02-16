//Package binarysearchtree 实现了二叉搜索树数据结构
package binarysearchtree

import (
	"container/list"
	"reflect"
)

//Item：树节点的数据可以是任意类型
type Item interface{}

//TreeNode：结点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//将传入数组转换成二叉树
func arrayToBinTree(nums []Item) *TreeNode {
	queue := list.New()
	if reflect.TypeOf(nums[0]) == nil {
		return nil
	}
	root := &TreeNode{nums[0].(int), nil, nil}
	queue.PushBack(root)
	for i := 1; queue.Len() != 0; {
		root := queue.Remove(queue.Front()).(*TreeNode)
		if i < len(nums) {
			if reflect.TypeOf(nums[i]) != nil {
				root.Left = &TreeNode{nums[i].(int), nil, nil}
				queue.PushBack(root.Left)
			}
			i++
		}
		if i < len(nums) {
			if reflect.TypeOf(nums[i]) != nil {
				root.Right = &TreeNode{nums[i].(int), nil, nil}
				queue.PushBack(root.Right)
			}
			i++
		}
	}
	return root
}

// 层次遍历
func levelorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		root = queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, root.Val)
		if root.Left != nil {
			queue.PushBack(root.Left)
		}
		if root.Right != nil {
			queue.PushBack(root.Right)
		}

	}
	return result
}

/* // 二叉树搜索：递归解法
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if val < root.Val {
        return searchBST(root.Left, val)
    }else if val > root.Val {
        return searchBST(root.Right, val)
    }else{
        return root
    }
} */

//二叉树搜索：迭代解法
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
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

//搜索最大结点
func findMax(root *TreeNode) *TreeNode {
	if root != nil {
		for root.Right != nil {
			root = root.Right
		}
	}
	return root
}

//二叉树插入
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{val, nil, nil}
	} else {
		if val > root.Val {
			root.Right = insertIntoBST(root.Right, val)
		} else if val < root.Val {
			root.Left = insertIntoBST(root.Left, val)
		}
	}
	return root
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
	return root
}
