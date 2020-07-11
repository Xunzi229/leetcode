//https://leetcode-cn.com/problems/sort-an-array/

package main

import "sort"

func sortArray(nums []int) []int {
    sort.Ints(nums)
    return nums
}