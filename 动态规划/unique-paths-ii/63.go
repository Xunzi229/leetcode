//https://leetcode-cn.com/problems/unique-paths-ii/
package unique_paths_ii

//[[0,0,0],[1,1,0],[0,0,0]]
// [[1]]
// [[0,1]]
// [[0,0],[0,1]]
// [[0,1],[1,0]]

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	// 2、初始化
	for i := 0; i <= len(obstacleGrid)-1; i++ {
		for j := 0; j <= len(obstacleGrid[i])-1; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, len(obstacleGrid[i]))
			}
			dp[i][j] = 1
		}
	}

	// 动态规划的过程
	for i := 0; i <= len(obstacleGrid)-1; i++ {
		for j := 0; j <= len(obstacleGrid[i])-1; j++ {
			// 当自身是 1 的时候这个时候是走不通的
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

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
