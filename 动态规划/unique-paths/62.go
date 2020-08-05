//https://leetcode-cn.com/problems/unique-paths/
package unique_paths

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	// 2、初始化
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n)
			}
			dp[i][j] = 1
		}
	}

	// 动态规划的过程
	for i := 0; i <= m-1; i++ {
		for j := 0; j <= n-1; j++ {
			if i != 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j]
			}
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}
