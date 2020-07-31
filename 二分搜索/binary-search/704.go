// https://leetcode-cn.com/problems/binary-search/submissions/
package binary_search

// [2,5] 5
// [] 2
// [2] 2
// [-1,0,3,5,9,12] 2
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
			continue
		}
		if nums[mid] < target {
			start = mid + 1
			continue
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

func searchDeep(num []int, target int, start, end int) int {
	if len(num) == 0 {
		return -1
	}
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if num[mid] == target {
		return mid
	}
	if num[mid] > target {
		return searchDeep(num, target, start, mid-1)
	}
	return searchDeep(num, target, start+1, mid)
}
