//https://leetcode-cn.com/problems/merge-intervals/
package main

import (
    "fmt"
)

func main()  {
    intervals := [][]int{
        {1,4},
        {0,0},
        {3,0},
        {5,0},
        {2,0},
    }
    r := merge(intervals)
    fmt.Println(r)
}
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {return nil}

    l := len(intervals)
    for i := 0; i <= l - 1; i++ {
        for j := i + 1; j <= l - i - 1; j++ {
            if intervals[j][0] < intervals[j-1][0] {
                intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
            }
        }
    }

    var tmpArray [][]int
    left, right := intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if right >= intervals[i][0] {
            if right < intervals[i][1] {
                right = intervals[i][1]
            }
        } else {
            tmpArray = append(tmpArray, []int{left, right})
            left, right = intervals[i][0], intervals[i][1]
        }
    }
    tmpArray = append(tmpArray, []int{left, right})
    return tmpArray
}