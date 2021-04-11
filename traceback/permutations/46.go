/*
https://leetcode-cn.com/problems/permutations/
*/
package permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)

	var backtrace func(n int, num []int)

	backtrace = func(n int, num []int) {
		if len(num) == len(nums) {
			tmp := make([]int, len(num))
			copy(tmp, num)
			result = append(result, tmp)
			return
		}

		for i := n; i < len(nums); i++ {
			exist := false
			for _, v := range num {
				if v == nums[i] {
					exist = true
					break
				}
			}
			if exist {
				continue
			}

			num = append(num, nums[i])
			backtrace(n, num)
			num = num[:len(num)-1]
		}
	}
	backtrace(0, []int{})
	return result
}
