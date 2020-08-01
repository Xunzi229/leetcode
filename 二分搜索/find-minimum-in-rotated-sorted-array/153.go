/*https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/*/
package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	if len(nums) == 0 {
        return -1
    }

    start := 0
    end := len(nums) - 1

    for start+1 < end {
        mid := start + (end-start)/2
        if nums[mid] <= nums[end] {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] > nums[end] {
        return nums[end]
    }
    return nums[start]
}
