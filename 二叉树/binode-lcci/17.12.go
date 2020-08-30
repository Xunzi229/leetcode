/*
https://leetcode-cn.com/problems/binode-lcci/
*/
package binode_lcci

import "fmt"

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

// 中序遍历
func convertBiNode(root *TreeNode) *TreeNode {
	var head *TreeNode
	upNode := new(TreeNode)

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node.Val)
		if head == nil {
			head = node
			upNode = node
		} else {
			upNode.Right = node
			upNode.Left = nil
			upNode = node
		}
		root = node.Right
	}

	return head
}
