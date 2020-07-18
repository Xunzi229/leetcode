//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 此题和https://leetcode-cn.com/problems/binary-tree-level-order-traversal/submissions/ 是同一个原理，
// 只是在存储的时候将结果放在前面就行
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result = make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)
	for len(stack) != 0 {
		tmpStack := make([]*TreeNode, 0)
		tmpR := make([]int, 0)

		for _, v := range stack {
			if v.Left != nil {
				tmpStack = append(tmpStack, v.Left)
			}
			if v.Right != nil {
				tmpStack = append(tmpStack, v.Right)
			}
			tmpR = append(tmpR, v.Val)
		}
		stack = tmpStack
		result = append([][]int{tmpR}, result...)
	}
	return result
}
