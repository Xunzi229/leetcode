// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

package largest_rectangle_in_histogram

//
func largestRectangleArea(heights []int) int {
	max := 0
	lenHeight := len(heights)

	for i, _ := range heights {
		largestRectangleAreaDeep(heights, i, i, lenHeight, &max)
	}
	return max
}

func largestRectangleAreaDeep(heights []int, currentIndex, rightIndex, lenHeight int, max *int) {
	minInt := 1<<63 - 1
	for {
		if rightIndex >= lenHeight {
			return
		}
		if heights[rightIndex] < minInt {
			minInt = heights[rightIndex]
		}
		left := minInt * (rightIndex - currentIndex + 1)
		if left > *(max) {
			*(max) = left
		}
		rightIndex++
	}
}

/*
使用递归， 超出时间限制， 所以这个方式并不合适
func largestRectangleArea(heights []int) int {
	max := 0
	lenHeight := len(heights)

	for i, _ := range heights {
		max = largestRectangleAreaDeep(heights, i, i, lenHeight, max)
	}
	return max
}

func largestRectangleAreaDeep(heights []int, currentIndex, rightIndex, lenHeight int, max int) int {
	if currentIndex > lenHeight-1 || rightIndex > lenHeight-1 {
		return 0
	}

	nums := append([]int{}, heights[currentIndex:rightIndex+1]...)
	sort.Ints(nums)
	left := nums[0] * (rightIndex - currentIndex + 1)
	right := largestRectangleAreaDeep(heights, currentIndex, rightIndex+1, lenHeight, max)

    tmpMax := 0
    if left >= right {
		tmpMax = left
	} else {
		tmpMax = right
	}
    if tmpMax > max {
        max = tmpMax
    }

	return max
}
*/
