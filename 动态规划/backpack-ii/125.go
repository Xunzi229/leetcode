/*
https://www.lintcode.com/problem/backpack-ii/description
*/
package backpack_ii

/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @param V: Given n items with value V[i]
 * @return: The maximum value
 */

/*
有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值. 问最多能装入背包的总价值是多大?
*/

// 求的结果是 总价值, 求的状态是 价值的变化
// 如何使用动态规划去做
func backPackII(m int, A []int, V []int) int {
	dp := make([][]int, len(A)+1)

	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)

	}

	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-A[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}

	return dp[len(A)][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
