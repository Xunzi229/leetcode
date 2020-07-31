//https://www.lintcode.com/problem/search-for-a-range/description
package search_for_a_range

/**
 * @param A: an integer sorted array
 * @param target: an integer to be inserted
 * @return: a list of length 2, [index1, index2]
 */
func searchRange(A []int, target int) []int {
	start, end := 0, len(A)-1

	startIndex, endIndex := -1, -1
	for start <= end && end >= 0 {
		mid := (start + end) / 2

		if A[mid] == target {
			startIndex, endIndex = mid, mid
			for startIndex > 0 {
				if A[startIndex-1] < target {
					break
				}
				startIndex = startIndex - 1
			}

			for endIndex < len(A)-1 {
				if A[endIndex+1] > target {
					break
				}
				endIndex = endIndex + 1
			}
			break
		}
		if A[mid] > target {
			end = mid - 1
		}
		if A[mid] < target {
			start = mid + 1
		}
	}

	return []int{startIndex, endIndex}
}
