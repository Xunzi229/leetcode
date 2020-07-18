//https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 考虑到是先访问左子树还是右子树
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [3,9,20,null,null,15,7]
// [1,2,3,4,null,null,5]
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result = make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	reverse := true
	for len(stack) != 0 {
		tmpStack := make([]*TreeNode, 0)
		tmpR := make([]int, 0)

		for _, v := range stack {
			if reverse {
				if v.Left != nil {
					tmpStack = append([]*TreeNode{v.Left}, tmpStack...)
				}
				if v.Right != nil {
					tmpStack = append([]*TreeNode{v.Right}, tmpStack...)
				}
			} else {
				if v.Right != nil {
					tmpStack = append([]*TreeNode{v.Right}, tmpStack...)
				}
				if v.Left != nil {
					tmpStack = append([]*TreeNode{v.Left}, tmpStack...)
				}
			}

			tmpR = append(tmpR, v.Val)

		}
		reverse = !reverse
		stack = tmpStack
		result = append(result, tmpR)
	}
	return result
}
