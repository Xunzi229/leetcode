//https://leetcode-cn.com/problems/validate-binary-search-tree/
package main

import (
	"fmt"
	"math"
)

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
func main(){
	fmt.Println(math.MinInt32, math.MaxInt32)
}
//[1, 1]
// [1,null,1]
// [10,5,15,null,null,6,20]
//
func isValidBST(root *TreeNode) bool {
	var result = make([]int, 0)

	inOrder(root, &result)
	for i:= 0; i < len(result) - 1; i ++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

// 中序遍历， 保证当前的值是单调递增的
func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// 此种方法单调递归， 保证大于最大值， 小于最小值
func isBT(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}
	if root.Val <= l || root.Val >= r {
		return false
	}

	isLeft := isBT(root.Left, l, root.Val)
	isRight := isBT(root.Right, root.Val, r)

	return isLeft && isRight
}