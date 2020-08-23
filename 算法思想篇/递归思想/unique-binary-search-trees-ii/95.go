/*
https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
*/
package unique_binary_search_trees_ii


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
     if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    var allTrees []*TreeNode
    for i := start; i <= end; i++ {
        leftTrees := helper(start, i - 1)
        rightTrees := helper(i + 1, end)
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                currTree := &TreeNode{i, nil, nil}
                currTree.Left = left
                currTree.Right = right
                allTrees = append(allTrees, currTree)
            }
        }
    }
    return allTrees
}