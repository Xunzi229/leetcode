/*
https://leetcode-cn.com/problems/combination-sum-iii/
*/

package combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	if n <= 0 {
		return nil
	}

	var result [][]int

	var backtrace func(c int, sum *int, num *[]int)

	backtrace = func(c int, sum *int, num *[]int) {
		if *sum == n && len(*num) == k {
			newNum := make([]int, len(*num))
			copy(newNum, *num)
			result = append(result, newNum)
			return
		}

		for i := c; i <= 9; i++ {
			if len(*num) == k {
				return
			}
			*num = append(*num, i)
			*sum += i
			backtrace(i+1, sum, num)
			*sum -= i
			*num = (*num)[:len(*num)-1]
		}
	}
	x := 0
	backtrace(1, &x, &[]int{})

	return result
}
