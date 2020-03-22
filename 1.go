package main

import (
    "fmt"
)

func main()  {
    var a = []int{3, 2, 4, 5}
    b := twoSum(a, 8)
    fmt.Println(b)
}
func twoSum(nums []int, target int) []int {
    for i, _ := range nums {
        for j, v := range nums[i+1:] {
            if v + nums[i] == target {
                return []int{i, j + i + 1}
            }
        }
    }
    return []int{}
}
