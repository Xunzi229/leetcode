// https://leetcode-cn.com/problems/number-of-islands/
package number_of_islands

// 使用标记法 ， 如果是已经是遍历过了，且符合对应的数 则改变值 如果没有遍历过 则不改对应值

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var max int = 49
	lenX := len(grid)
	lexY := len(grid[0])

	for _, v := range grid {
		fmt.Println(v)
	}
	fmt.Println("---")

	//
	for x, _ := range grid {
		for y, v2 := range grid[x] {
			if string(v2) == "1" {
				max += 1
				numIslandsDeep(&grid, x, y, max, lenX, lexY)
			}
		}
	}

	for _, v := range grid {
		fmt.Println(v)
	}

	return max - 49
}

func numIslandsDeep(grid *[][]byte, x, y, max int, lenx, leny int) {
	(*grid)[x][y] = byte(max)

	if x-1 >= 0 && string((*grid)[x-1][y]) == "1" {
		numIslandsDeep(grid, x-1, y, max, lenx, leny)
	}
	if y-1 >= 0 && string((*grid)[x][y-1]) == "1" {
		numIslandsDeep(grid, x, y-1, max, lenx, leny)
	}

	if x+1 < lenx && string((*grid)[x+1][y]) == "1" {
		numIslandsDeep(grid, x+1, y, max, lenx, leny)
	}

	if y+1 < leny && string((*grid)[x][y+1]) == "1" {
		numIslandsDeep(grid, x, y+1, max, lenx, leny)
	}
}
