/*
https://leetcode-cn.com/problems/subsets-ii/
*/
package subsets_ii

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	sort.Ints(nums)

	result := make([][]int, 0)
	result = append(result, []int{})
	var backtrace func(n int, num []int)

	backtrace = func(n int, num []int) {
		for i := n; i < len(nums); i++ {
			num = append(num, nums[i])
			t := false
			for _, nm := range result {
				if equal(&nm, &num) {
					t = true
					break
				}
			}
			if !t {
				newNum := append([]int{}, num...)
				result = append(result, newNum)
			}
			backtrace(i+1, num)
			num = num[:len(num)-1]
		}
	}
	backtrace(0, []int{})
	return result
}

func equal(a, b *[]int) bool {
	if len(*a) != len(*b) {
		return false
	}

	for i := 0; i < len(*a); i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}
