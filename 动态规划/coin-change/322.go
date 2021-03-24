/*
https://leetcode-cn.com/problems/coin-change/
*/
/**/
package coin_change

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
*/

// 凑零钱
// 动态规划
// 如何找状态转移公式

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // 存 0
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// 如何找状态转移方程
	//
	for i := 1; i <= amount; i++ {
		for j := 0; j <= len(coins)-1; j++ {
			if i-coins[j] >= 0 { // 保证 加一个的时候不会超出当前的钱
				dp[i] = min(dp[i], dp[i-coins[j]]+1) // 这个地方 保证上一个最少的 + 1, 相比较
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
