//https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/

package search_in_rotated_sorted_array

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start, end := 0, len(nums)-1
	for start+1 < end {

		for start+1 <= end && nums[start] == nums[start+1] {
			start += 1
		}
		for end-1 >= 0 && end-1 >= start && nums[end] == nums[end-1] {
			end -= 1
		}

		mid := start + (end-start)/2
		if nums[end] == target || nums[start] == target || nums[mid] == target {
			return true
		}

		if nums[end] == nums[start] {
			end -= 1
			if nums[end] == target {
				return true
			}
			for end-1 >= 0 && end-1 >= start && nums[end] == nums[end-1] {
				end -= 1
			}
		}

		if nums[start] < nums[end] {
			if nums[mid] > target {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] > nums[end] {
				if nums[mid] > target && target > nums[end] {
					end = mid
					continue
				}
				if nums[mid] > target && target < nums[end] {
					start = mid
					continue
				}
				if nums[mid] < target {
					start = mid
				}
			} else {
				if nums[mid] > target {
					end = mid
				} else {
					if nums[end] > target {
						start = mid
					} else {
						end = mid
					}
				}
			}
		}
	}

	return nums[start] == target || nums[end] == target
}
