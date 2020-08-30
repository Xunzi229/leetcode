/*
https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/
*/
package maximum_sum_bst_in_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
二叉搜索树的定义如下：

任意节点的左子树中的键值都小于此节点的键值。
任意节点的右子树中的键值都 大于此节点的键值。
任意节点的左子树和右子树都是二叉搜索树。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
[1,null,10,-5,20]
[0,9,-8,6,-6,9,3,-5,1,7,1,0,null,-6,null,-4,1,null,3,2,null,null,null,null,null,null,null,null,null,null,null,4,null,10,8,null,null,null,null,1,13,-1,2,10,16,null,null,null,6,null,12,null,17]

*/
func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := 0
	_, _, _, _ = isBST(root, &maxValue)
	return maxValue
}

// 入参, 节点 和最大值
// 出参  是否是二叉搜索树, 上一层值得总和
/*
保证每一层都是二叉搜索树
使用递归
*/
func isBST(root *TreeNode, maxValue *int) (bool, int, int, int) {
	if root.Left == nil && root.Right == nil {
		if *maxValue < root.Val {
			*maxValue = root.Val
		}
		return true, root.Val, root.Val, root.Val
	}
	if root.Left != nil && root.Right == nil {
		ok1, value1, max, min := isBST(root.Left, maxValue)

		if ok1 && root.Val > max {
			if *maxValue < root.Val+value1 {
				*maxValue = root.Val + value1
			}
		}
		return ok1 && root.Val > max, root.Val + value1, Max(root.Val, max), Min(root.Val, min)
	}
	if root.Left == nil && root.Right != nil {
		ok2, value2, max, min := isBST(root.Right, maxValue)
		if ok2 && root.Val < min {
			if *maxValue < root.Val+value2 {
				*maxValue = root.Val + value2
			}
		}
		return ok2 && root.Val < min, root.Val + value2, Max(root.Val, max), Min(root.Val, min)
	}

	ok1, value1, max1, min1 := isBST(root.Left, maxValue)
	ok2, value2, max2, min2 := isBST(root.Right, maxValue)
	if ok1 && ok2 && root.Val > max1 && root.Val < min2 {
		if *maxValue < value2+value1+root.Val {
			*maxValue = value2 + value1 + root.Val
		}
		return true, value2 + value1 + root.Val, Max(Max(root.Val, max1), max2), Min(Min(root.Val, min1), min2)
	}
	return false, value2 + value1 + root.Val, Max(Max(root.Val, max1), max2), Min(Min(root.Val, min1), min2)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
