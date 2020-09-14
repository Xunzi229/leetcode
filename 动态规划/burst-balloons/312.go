/*
https://leetcode-cn.com/problems/burst-balloons/
*/
package burst_balloons

func maxCoins(nums []int) int {
	// 收尾增加边界
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	// Dp函数
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := 1; i <= len(nums)-2; i++ {
		nums[i] = nums[i-1]
	}

	for i := 3; i < len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			res := 0
			for k := j + 1; k < j+i-1; k++ {
				left := dp[i][k]
				right := dp[k][j+i-1]
				res = max(res, left+nums[j]*nums[k]*nums[j+i-1]+right)
			}
			dp[i][j+i-1] = res
		}
	}
	return dp[0][len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
