//https://leetcode-cn.com/problems/sort-colors/
package sort_colors

func sortColors(nums []int) {
	center := 0
	notRemove := 0
	for i, v := range nums {
		if v != 2 {
			if i-notRemove >= 0 {
				nums[i-notRemove], nums[i] = nums[i], nums[i-notRemove]
				nums[i-notRemove], nums[center] = nums[center], nums[i-notRemove]
				if v == 0 {
					center++
				}
			}
		} else {
			notRemove++
		}
	}
}
