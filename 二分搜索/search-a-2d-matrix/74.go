//https://leetcode-cn.com/problems/search-a-2d-matrix/
package search_a_2d_matrix

// [[1],[3],[5]]
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	iStart, iEnd := 0, len(matrix)-1
	for iStart+1 < iEnd {
		iMid := iStart + (iEnd-iStart)/2

		if len(matrix[iStart]) > 0 {
			if matrix[iStart][0] <= target && matrix[iStart][len(matrix[iStart])-1] >= target {
				nums := matrix[iStart]
				//此处判断 target 是否在 nums 中就行
				return searchMatrixDeep(nums, target)
			}
		}
		//
		if len(matrix[iEnd]) > 0 {
			if matrix[iEnd][0] <= target && matrix[iEnd][len(matrix[iEnd])-1] >= target {
				nums := matrix[iEnd]
				//此处判断 target 是否在 nums 中就行
				return searchMatrixDeep(nums, target)
			}
		}
		//
		if len(matrix[iMid]) > 0 {
			if matrix[iMid][0] > target {
				iEnd = iMid - 1
				continue
			}
			if iMid < target {
				iStart = iMid + 1
				continue
			}
			return searchMatrixDeep(matrix[iMid], target)
		} else {
			// 向下查找
			for {
				iMid -= 1
				if iMid == 0 {
					break
				}
				if len(matrix[iMid]) > 0 {
					break
				}
			}
			if matrix[iMid][len(matrix[iMid])-1] < target {
				for {
					iMid += 1
					if iMid == len(matrix) {
						break
					}
					if len(matrix[iMid]) > 0 {
						break
					}
				}

				iStart = iMid
				continue
			}
			iEnd = iMid
		}

	}
	return false
}

// 二分搜索
func searchMatrixDeep(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	var start, end = 0, len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}
