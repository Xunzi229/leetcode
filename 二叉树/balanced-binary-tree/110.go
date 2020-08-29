/*
https://leetcode-cn.com/problems/balanced-binary-tree/
*/
package balanced_binary_tree

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

/*
[1,2,2,3,null,null,3,4,null,null,4]
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok1, _ := deepIsBalanced(root)
	return ok1
}

func deepIsBalanced(root *TreeNode) (bool, int) {
	if root.Left == nil && root.Right == nil {
		return true, 0
	}
	if root.Left != nil && root.Right == nil {
		ok, deep := deepIsBalanced(root.Left)
		return ok && deep < 1, deep + 1
	}
	if root.Right != nil && root.Left == nil {
		ok, deep := deepIsBalanced(root.Right)

		return ok && deep < 1, deep + 1
	}
	ok1, d1 := deepIsBalanced(root.Left)
	ok2, d2 := deepIsBalanced(root.Right)

	if d1 > d2 {
		return ok1 && ok2 && d1-d2 <= 1, d1 + 1
	}
	return ok1 && ok2 && d2-d1 <= 1, d2 + 1
}
