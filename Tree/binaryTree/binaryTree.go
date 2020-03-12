package binarytree

import (
	"container/list"
	"reflect"
)

type Item interface{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// 前序遍历：递归解法
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, root.Val)
		result = append(result, preorderTraversal(root.Left)...)
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}

/* // 前序遍历：迭代解法1
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			result = append(result, root.Val)
			root = root.Left
		}
		if stack.Len() != 0 {
			root = stack.Remove(stack.Back()).(*TreeNode)
			root = root.Right
		}
	}
	return result
} */

/* // 前序遍历：迭代解法2
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		tmp := stack.Remove(stack.Back()).(*TreeNode)
		if tmp != nil {
			result = append(result, tmp.Val)
			stack.PushBack(tmp.Right)
			stack.PushBack(tmp.Left)
		}
	}
	return result
} */

// 中序遍历：递归解法
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

/* // 中序遍历：迭代解法
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() != 0 {
			root = stack.Remove(stack.Back()).(*TreeNode)
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
} */

/* // 后序遍历：递归解法
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, postorderTraversal(root.Left)...)
		result = append(result, postorderTraversal(root.Right)...)
		result = append(result, root.Val)
	}
	return result
} */

// 后序遍历：迭代解法
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := list.New()
	tag := make(map[*TreeNode]bool)
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() != 0 {
			root = stack.Back().Value.(*TreeNode)
			if !tag[root] {
				tag[root] = true
				root = root.Right
			} else {
				root = stack.Remove(stack.Back()).(*TreeNode)
				result = append(result, root.Val)
				root = nil
			}
		}
	}
	return result
}

// 层次遍历
func levelorderTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		currentLevel := []int{}
		currentLevelLen := queue.Len()
		for i := 0; i < currentLevelLen; i++ {
			root = queue.Remove(queue.Front()).(*TreeNode)
			currentLevel = append(currentLevel, root.Val)
			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}
