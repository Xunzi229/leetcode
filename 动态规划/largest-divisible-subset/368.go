/*
https://leetcode-cn.com/problems/largest-divisible-subset/
*/
package largest_divisible_subset

import (
	"sort"
)

/*
给出一个由无重复的正整数组成的集合，找出其中最大的整除子集，
子集中任意一对 (Si，Sj) 都要满足：Si % Sj = 0 或 Sj % Si = 0。
如果有多个目标子集，返回其中任何一个均可。
[5,9,18,54,108,540,90,180,360,720]
*/

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	// 对数据进行排序
	sort.Ints(nums)

	//
	dp := make([]int, n)

	maxValue := 1
	maxIndex := 0

	// 初始化 为 1
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if maxValue < dp[i] {
			maxIndex = i
			maxValue = dp[i]
		}
	}

	collection := make([]int, 0)
	for i := maxIndex; i >= 0; i-- {
		if (nums[maxIndex]%nums[i] == 0) && (dp[i] == maxValue) {
			collection = append(collection, nums[i])
			maxValue--
			maxIndex = i
		}
	}
	return collection
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
