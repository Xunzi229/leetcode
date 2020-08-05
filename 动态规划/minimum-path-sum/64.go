//https://leetcode-cn.com/problems/minimum-path-sum/
package minimum_path_sum

import "fmt"

func minPathSum(grid [][]int) int {
	c := len(grid)
	dp := make([][]int, c)
	// 2、初始化
	for i := 0; i < c; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if dp[i] == nil {
				dp[i] = make([]int, len(grid[i]))
			}
			dp[i][j] = grid[i][j]
		}
	}

	// 动态规划的过程
	for i := 0; i <= len(grid)-1; i++ {
		for j := 0; j <= len(grid[i])-1; j++ {
			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
			if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	fmt.Println(len(dp[len(grid)-1]))
	return dp[len(grid)-1][len(dp[len(grid)-1])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
