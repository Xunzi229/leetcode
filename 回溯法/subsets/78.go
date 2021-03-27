/*
https://leetcode-cn.com/problems/subsets/
*/
package subsets

/*
给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

//
/*
1. 首先 是结果可计算出来 2^len(nums)次方
2.
*/
func subsets(nums []int) [][]int {
	set := make([][]int, 0)

	var backtrace func(ps, num int, ts []int)

	backtrace = func(ps, num int, ts []int) {
		if num == len(ts) {
			tmp := append([]int{}, ts...)
			set = append(set, tmp)
			return
		}

		for i := ps; i < len(nums); i++ {
			ts = append(ts, nums[i])
			backtrace(i+1, num, ts)
			ts = ts[:len(ts)-1]
		}
	}

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		backtrace(0, i, cur)
	}
	return set
}
