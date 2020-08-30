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
	nums := make([]*TreeNode, 0)
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		nums = append(nums, node)
		root = node.Right
		node.Left = nil
		node.Right = nil
	}
	if len(nums) <= 1 {
		return root
	}

	center := len(nums) / 2
	root = nums[center]

	numBalance(nums[center], 0, center-1, &nums)
	numBalance(nums[center], center+1, len(nums)-1, &nums)
	return root
}

func numBalance(root *TreeNode, start, end int, nums *[]*TreeNode) {
	if end < 0 || end >= len(*nums) {
		return
	}

	center := (end - start) / 2
	if (*nums)[center].Val > root.Val {
		root.Left = (*nums)[center]
	}
	if (*nums)[center].Val < root.Val {
		root.Left = (*nums)[center]
	}
	numBalance((*nums)[center], 0, center-1, nums)
	numBalance((*nums)[center], center+1, len(*nums)-1, nums)
}
