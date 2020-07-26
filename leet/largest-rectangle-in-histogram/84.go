// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

package largest_rectangle_in_histogram

// TODO next
func largestRectangleArea(heights []int) int {
	lenHeight := len(heights)
	if lenHeight == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == lenHeight {
			cur = 0
		} else {
			cur = heights[i]
		}
		// 当前高度小于栈，则将栈内元素都弹出计算面积
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]

			// 计算宽度
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = Max(max, h*w)
		}
		// 记录索引即可获取对应元素
		stack = append(stack, i)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 该方法虽然执行的速度上有所提升并且通过测试但是还未是效率上最好的
/*
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
		if heights[rightIndex] < heights[currentIndex] {
			if heights[rightIndex]+lenHeight-currentIndex+1 < *max {
				return
			}
		}

		rightIndex++
	}
}
*/

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
