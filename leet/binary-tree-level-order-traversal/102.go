// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

package main

func main() {

}

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

func levelOrder(root *TreeNode) [][]int {
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
		result = append(result, tmpR)
	}
	return result
}
