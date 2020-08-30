/*
https://leetcode-cn.com/problems/balance-a-binary-search-tree/
*/
package balance_a_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	if root == nil {
		return root
	}

	// 使用中序遍历 获取所有的元素值
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, node.Val)
		root = node.Right
	}
	if len(nums) <= 1 {
		return root
	}
	return numBalance(0, len(nums)-1, &nums)
}

// 使用递归
func numBalance(start, end int, nums *[]int) *TreeNode {
	if start > end {
		return nil
	}

	center := (start + end) / 2
	root := &TreeNode{Val: (*nums)[center]}
	root.Left = numBalance(start, center-1, nums)
	root.Right = numBalance(center+1, end, nums)
	return root
}
