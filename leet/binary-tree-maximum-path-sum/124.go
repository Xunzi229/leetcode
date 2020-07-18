// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 开始节点 到终止的节点 不一定经过根节点
// 清楚的理解分治， 这里最大值其实不一定经过根节点， 那么有可能的是 在某个节点其实已经达到最大值
// 所以在某个节点的时候 ， 我们需要存的是左右节点获得的最大值， 和比较加上其节点获得的值比较

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ResultType struct {
	SinglePath int // 保存单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func main() {
}

func maxPathSum(root *TreeNode) int {
	childVal := helper(root)
	return childVal.MaxPath
}

func helper(root *TreeNode) ResultType {
	// check
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}
	// Divide
	left := helper(root.Left)
	right := helper(root.Right)

	// Conquer
	result := ResultType{}

	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}

	// 求两边加根最大值
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
