/*
https://leetcode-cn.com/problems/combination-sum-ii/
*/

package combination_sum_ii

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)

	result := make([][]int, 0)
	var backtrace = func(n, sum int, num []int) {}

	backtrace = func(n, sum int, num []int) {
		if n >= len(candidates) || candidates[n]+sum > target {
			return
		}
		for i := n; i < len(candidates); i++ {
			if candidates[i]+sum == target {
				newNum := append([]int{candidates[i]}, num...)
				if !isExist(&result, &newNum) {
					result = append(result, newNum)
				}
				return
			}

			sum += candidates[i]
			num = append(num, candidates[i])

			backtrace(i+1, sum, num)

			sum -= candidates[i]
			num = num[:len(num)-1]
		}
	}

	for i := 0; i < len(candidates); i++ {
		if i >= 1 && candidates[i] == candidates[i-1] {
			continue
		}
		backtrace(i, 0, []int{})
	}

	return result
}

func isExist(nums *[][]int, num *[]int) bool {
	f := func(a, b *[]int) bool {
		if (a == nil) != (b == nil) {
			return false
		}

		if len(*a) != len(*b) {
			return false
		}

		for i := range *a {
			if (*a)[i] != (*b)[i] {
				return false
			}
		}

		return true
	}
	for _, n := range *nums {
		if ok := f(&n, num); ok {
			return true
		}
	}
	return false
}
