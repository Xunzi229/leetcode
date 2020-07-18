//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 判断该节点在左子树 还是在 右子树上， 如果在做子树 则返回做子树的根节点， 如果在右子树 则返回右子树的根节点， 如果都不在， 则返回根节点
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}
