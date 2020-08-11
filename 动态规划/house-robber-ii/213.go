/*
https://leetcode-cn.com/problems/house-robber-ii/
*/
package house_robber_ii

/*
213. 打家劫舍 II
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
*/
// [1,2,1,1]
// [2,7,9,3,1]
// [1,1,3,6,7,10,7,1,8,5,9,1,4,4,3]
// [1,1,2,1]
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 0 ~ n-1
	dp1 := make([]int, len(nums))
	dp1[0] = nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	dp1[1] = max(nums[1], nums[0])
	for i := 2; i <= len(nums)-2; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}

	// 过滤只有两个元素的情况
	if len(nums) == 2 {
		return dp1[1]
	}

	// 过滤只有三个元素的情况
	if len(nums) == 3 {
		return max(dp1[1], nums[2])
	}

	// 1 ~ n
	dp2 := make([]int, len(nums))
	dp2[1] = nums[1]
	dp2[2] = max(nums[1], nums[2])
	for i := 3; i <= len(nums)-1; i++ {
		dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
	}

	return max(dp1[len(nums)-2], dp2[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
[1 1 4 7 11 17 18 18 26 26 35 35 39 39 39]
[true false true false true false true true true true true true true true false]
*/
