package main

import "fmt"

func main() {
	Queen()
}

// 在 8 * 8 的方格上, 每一行每一列 斜方向只能放一个皇后, 那么一共有多少种摆法
func Queen() {
	var nums = make([][]int, 8)
	for i := 0; i < 8; i++ {
		nums[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			nums[i][j] = 0
		}
	}

	queenDeep(nums, 0)
}

// 遍历
func queenDeep(nums [][]int, row int) {
	if row == len(nums) {
		// 打印棋盘
		for i, _ := range nums {
			for _, v := range nums[i] {
				fmt.Printf("%2d", v)
			}
			fmt.Printf("\n")
		}

		fmt.Printf("----------------------------\n")
		return
	}

	n := len(nums[row])
	for col := 0; col < n; col++ {

		// 排除不合法选择
		if !isValid(nums, row, col) {
			continue
		}

		// 做选择
		nums[row][col] = 1

		// 进入下一行决策
		queenDeep(nums, row+1)

		// 撤销选择
		nums[row][col] = 0
	}
}

// 用于判断该点能不能放置皇后
func isValid(nums [][]int, row, col int) bool {
	n := len(nums)

	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if nums[i][col] == 1 {
			return false
		}
	}

	// 检查右上方是否有皇后互相冲突
	i, j := row-1, col+1
	for j < n && i >= 0 {
		if nums[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	// 检查左上方是否有皇后互相冲突
	i, j = row-1, col-1
	for j >= 0 && i >= 0 {
		if nums[i][j] == 1 {
			return false
		}
		i--
		j--
	}
	return true
}
