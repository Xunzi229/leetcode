// https://leetcode-cn.com/problems/search-insert-position/
package search_insert_position

// [3,5,7,9,10]  8
// [1,3,5,6] 2
// [1,3,5,6] 5
// [1,3,5,6] 0
// [1,3,5,6] 7
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
			continue
		}
		start = mid
	}

	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		return end + 1
	}
	return 0
}
