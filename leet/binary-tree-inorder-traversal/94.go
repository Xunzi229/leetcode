//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/
package binary_tree_inorder_traversal

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

func inorderTraversal(root *TreeNode) []int {
	var values = make([]int, 0)
	if root == nil {
		return values
	}

	stacks := make([]*TreeNode, 0)
	for len(stacks) != 0 || root != nil {
		for root != nil {
			stacks = append(stacks, root)
			root = root.Left
		}

		root = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]
		values = append(values, root.Val)
		root = root.Right
	}

	return values
}
