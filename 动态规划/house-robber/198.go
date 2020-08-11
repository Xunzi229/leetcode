/*
https://leetcode-cn.com/problems/house-robber/
*/
package house_robber

/*
198. 打家劫舍
等级: 简单
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
*/

// 状态在过程中的变化
// 状态转移方程 dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	dp[1] = max(nums[1], nums[0])

	for i := 2; i <= len(nums)-1; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
