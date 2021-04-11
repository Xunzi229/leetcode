/*
https://leetcode-cn.com/problems/permutations-ii/
*/
package permutations_ii

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	et := map[int]int{}

	var backtrace func(num []int, spl map[int]int)

	for i := 0; i < len(nums); i++ {
		et[nums[i]]++
	}

	backtrace = func(num []int, spl map[int]int) {
		if len(num) == len(nums) {
			tmp := make([]int, len(num))
			copy(tmp, num)
			result = append(result, tmp)
			return
		}

		y := map[int]int{} // 保证当前层不再使用相同的数， 造成重复
		// spl // 确定重复的数还能使用的次数
		for i := 0; i < len(nums); i++ {
			if spl[nums[i]] <= 0 || y[nums[i]] > 0 {
				continue
			}

			num = append(num, nums[i])
			spl[nums[i]]--
			y[nums[i]]++
			backtrace(num, spl)
			num = num[:len(num)-1]
			spl[nums[i]]++
		}
	}
	backtrace([]int{}, et)
	return result
}
