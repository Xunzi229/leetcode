package maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+dp[i-1] >= nums[i] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
