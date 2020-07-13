// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode104 struct {
	Val   int
	Left  *TreeNode104
	Right *TreeNode104
}

func maxDepth(root *TreeNode104) int {
	var max = 0
	//tmpTraversalDepth(root, 0, &max)

	tmpTraversalDepth(root, &max)
	return max
}

// [3,9,20,null,null,15,7]
// [0]
// [1,2]
func tmpTraversalDepth(node *TreeNode104, max *int) {
	if node == nil {
		return
	}
	var queue []*TreeNode104
	queue = append(queue, node)

	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode104, 0)
		for _, v := range queue {
			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}
		queue = tmpQueue
		*max++
	}
	return
}

/*
使用递归
func tmpTraversalDepth(node *TreeNode104, i int, max *int) {
	if node == nil {
		return
	}
	i++
	if i > *max {
		*max = i
	}
	tmpTraversalDepth(node.Left, i, max)
	tmpTraversalDepth(node.Right, i, max)
}
*/
