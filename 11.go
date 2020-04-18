// https://leetcode-cn.com/problems/container-with-most-water/

package main

import "fmt"

func main()  {
    fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

// 方式1
/*
此种方式比较暴力， 这样需要全部遍历下来
*/
/*func maxArea(height []int) int {
    max := 0

    l := len(height)
    for i := 0; i <= l - 1; i++ {
        for j := i + 1; j <= l - 1; j++ {
            mi := height[i]
            if height[i] > height[j] {
               mi = height[j]
            }
            if (j - i) * mi > max {
                max = (j - i) * mi
            }
        }
    }
    return max
}
*/
// 方式2
// 需要了解的是 ， 盛水需要判断短的一端， ，所以短的一端向里增加一个位置，，如果长的一端向里增加则只可能更加面积少
func maxArea(height []int) int {
    max := 0

    left, right := 0, len(height) -1
    for {
        if left == right {
            break
        }

        if height[left] < height[right] {
            if (right - left) * height[left] > max {
                max = (right - left) * height[left]
            }
            left++
            continue
        }

        if height[left] >= height[right] {
            if (right - left) * height[right] > max {
                max = (right - left) * height[right]
            }
            right--
        }
    }
    return max
}