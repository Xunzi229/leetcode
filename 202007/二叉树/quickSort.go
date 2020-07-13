package main

import (
	"fmt"
)

func main() {
	var nums = []int{6, 1, 4, 2, 3, 7, 5}
	QuickSort(nums)
	fmt.Println(nums)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	// 防止超过界限
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			fmt.Println("i, j, p:", i, j, p, nums[i], nums[j], nums)
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值

	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
