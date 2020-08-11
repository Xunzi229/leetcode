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
*/

func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sort.Ints(nums)

	dp := make([][]int, len(nums))

	return dp[0]
}
