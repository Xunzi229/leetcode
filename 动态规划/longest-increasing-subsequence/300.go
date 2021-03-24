/*
https://leetcode-cn.com/problems/longest-increasing-subsequence/
*/
package longest_increasing_subsequence

// 动态规划
// 找的是整体的非连续的最长的子序列, 2, 3, 5, 4, 6  =>  2,3,4,6   长度为4
//
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 假定dp中的每个index 表示的nums对应的index前面最长的序列
	dp := make([]int, len(nums))
	var maxLen = 0
	for i := 0; i < len(nums); i++ {
		t := 0
		for j := 0; j < i; j++ {
			// dp[j-1] 其实已经知道他的最长的连续长度了
			// 只有当这个值j 对应的值还小于这个值的时候才会在之前的所以的最长的长度上加1
			if nums[i] > nums[j] {
				t = max(t, dp[j])
			}
		}
		dp[i] = t + 1
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
