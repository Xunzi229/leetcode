/*https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/*/
package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	var start, end = 0, len(nums) - 1

	tmpStart := nums[0]
	for start+1 < end {
		mid := start + (end-start)/2

		if nums[mid] > nums[mid+1] {
			if nums[mid+1] < tmpStart {
				return nums[mid+1]
			}
			return tmpStart
		}
		if mid-1 >= 0 {
			left := findMin(nums[start : mid-1])
			right := findMin(nums[mid+1 : end])
			if left > right {
				return right
			}
			return left
		}
		return findMin(nums[mid+1 : end])
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
