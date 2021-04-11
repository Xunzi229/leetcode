/*
https://leetcode-cn.com/problems/combinations/
*/

package combine

func combine(n int, k int) [][]int {
	if k <= 0 || n <= 0 {
		return nil
	}

	result := make([][]int, 0)

	var backtrace = func(c int, num []int) {}

	backtrace = func(c int, num []int) {
		if len(num) == k {
			newNum := make([]int, k)
			copy(newNum, num)
			result = append(result, newNum)
			return
		}

		for i := c; i <= n; i++ {
			num = append(num, i)
			backtrace(i+1, num)
			num = num[:len(num)-1]
		}
	}
	backtrace(1, []int{})
	return result
}
