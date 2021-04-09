/*
https://leetcode-cn.com/problems/combination-sum/
*/

package combination_sum

import "sort"

// 为啥数组可以用 nil 返回
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	// 排序后保证后面遍历不需要回头再加减

	sort.Ints(candidates)

	result := make([][]int, 0)

	var backtrace func(n, sum int, nums []int)

	backtrace = func(n, sum int, nums []int) {
		if n >= len(candidates) {
			return
		}
		for i := n; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				continue
			}
			if sum+candidates[i] == target {
				nums = append(nums, candidates[i])
				x := append([]int{}, nums...)
				result = append(result, x)
				return
			}
			sum += candidates[i]
			nums = append(nums, candidates[i])
			backtrace(i, sum, nums)
			sum -= candidates[i]
			nums = nums[:len(nums)-1]
		}
	}

	for i, v := range candidates {
		if v == target {
			result = append(result, []int{v})
			continue
		}
		backtrace(i, v, []int{candidates[i]})
	}

	return result
}
