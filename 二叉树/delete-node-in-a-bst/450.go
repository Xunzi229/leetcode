/*
https://leetcode-cn.com/problems/delete-node-in-a-bst/
*/
package delete_node_in_a_bst

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

// 如果想删除其中一个元素, 谦虚遍历就可以, 如果当前节点符合
// 将指向当前节点的
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 只有左边的节点, 替换为右边
	// 只右右边的节点 替换为左
	// 有左右节点 左子节点连接到右边最左节点即可
	//
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}
