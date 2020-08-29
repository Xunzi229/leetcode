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

// 平衡的二叉树, 使用递归比较好做, 在每一层都需要判断该层是不是平衡二叉树, 如果不是
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok1, _ := deepIsBalanced(root)
	return ok1
}

func deepIsBalanced(root *TreeNode) (bool, int) {
	/*  一  */
	//if root.Left == nil && root.Right == nil {
	//	return true, 0
	//}
	//if root.Left != nil && root.Right == nil {
	//	ok, deep := deepIsBalanced(root.Left)
	//	return ok && deep < 1, deep + 1
	//}
	//if root.Right != nil && root.Left == nil {
	//	ok, deep := deepIsBalanced(root.Right)
	//	return ok && deep < 1, deep + 1
	//}
	/*------------------*/
	/*  二  */
	if root == nil {
		return true, 0
	}
	/*------------------*/
	ok1, d1 := deepIsBalanced(root.Left)
	ok2, d2 := deepIsBalanced(root.Right)
	return ok1 && ok2 && abs(d1, d2) <= 1, max(d1, d2) + 1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
