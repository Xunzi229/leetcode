//https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/
package binary_search_tree_to_greater_sum_tree

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

func bstToGst(root *TreeNode) *TreeNode {
	var list = make([]int, 0)
	deepBstToGst(root, &list)
	return root
}

func deepBstToGst(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	deepBstToGst(root.Right, list)
	*list = append(*list, root.Val)
	root.Val = sumInt(list)
	deepBstToGst(root.Left, list)
}

func sumInt(list *[]int) int {
	if len(*list) == 0 {
		return 0
	}

	var s int
	for _, v := range *list {
		s += v
	}
	return s
}
