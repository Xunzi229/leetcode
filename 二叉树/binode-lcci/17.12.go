/*
https://leetcode-cn.com/problems/binode-lcci/
*/
package binode_lcci

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
/*
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
*/

func convertBiNode(root *TreeNode) *TreeNode {
	header := &TreeNode{}
	node := header

	deepConvertBiNode(root, node)

	return header.Right
}

// 进行中序遍历
func deepConvertBiNode(root, node *TreeNode) *TreeNode {
	if root != nil {
		node = deepConvertBiNode(root.Left, node)
		root.Left = nil
		node.Right = root
		node = root
		node = deepConvertBiNode(root.Right, node)
	}
	return node
}
