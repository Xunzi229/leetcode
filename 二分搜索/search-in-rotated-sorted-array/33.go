// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/submissions/
package search_in_rotated_sorted_array

// 有序数组进行的旋转
// 最后一个

// 判断这个 target是在左边 还是 右边，
// 如果 tagert 是大于最后一个值 在左边
// 如果 target 是小于最后一个值 在右边
// [4,5,6,7,8,1,2,3] 8
// [1] 1
// [4,5,6,7,0,1,2] 0
// [4,5,6,7,0,1,2] 6
// [1,3] 2

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start+1 <= end {
		mid := start + (end-start)/2
		if nums[end] == target {
			return end
		}
		if nums[start] == target {
			return start
		}
		if nums[mid] == target {
			return mid
		}
		if nums[start] < nums[end] {
			if nums[mid] > target {
				end = mid
			} else {
				start = mid
			}
		} else {
			// nums[start] > nums[end]
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

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}

	return -1
}
