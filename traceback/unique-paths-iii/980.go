/*
https://leetcode-cn.com/problems/unique-paths-iii/
*/
package unique_paths_iii

/*
1. 遍历获取起始方格、结束方格的位置、和总的需要过的方格数
[
[1,  0,  0,  0]
[0,  0,  0,  0]
[0,  0,  2,  -1]
]
*/

func uniquePathsIII(grid [][]int) int {
	// x, y // x 表示多少行， y表示每行多少数据

	// 向上 x--
	// 向下 x++
	// 向左 y--
	// 向右 y++
	start := []int{0, 0}
	end := []int{0, 0}
	needStep := 0

	var repeatGrid = make([][]int, len(grid))
	for x := 0; x < len(grid); x++ {
		repeatGrid[x] = make([]int, len(grid[x]))
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				start = []int{x, y}
				repeatGrid[x][y] = -1
			}
			if grid[x][y] == 2 {
				end = []int{x, y}
			}
			if grid[x][y] == 0 {
				needStep++
			}

			if grid[x][y] == -1 {
				repeatGrid[x][y] = -1
			}
		}
	}

	var res int

	var backtrace func(x, y, step int, colorGrid *[][]int)

	backtrace = func(x, y, step int, colorGrid *[][]int) {
		// 到达终点的标志, 为啥等于 -1, 因为 0 的表示都走完了，最后一步到这个地方就变成-1了
		if x == end[0] && y == end[1] {
			if step == -1 {
				res++
			}
			return
		}
		if step == -1 {
			return
		}

		for {
			// 向上
			if x-1 >= 0 && grid[x-1][y] != -1 && repeatGrid[x-1][y] != -1 {
				repeatGrid[x-1][y] = -1
				step--
				backtrace(x-1, y, step, &repeatGrid)
				step++
				repeatGrid[x-1][y] = 0
			}

			// 向下
			if x+1 < len(grid) && grid[x+1][y] != -1 && repeatGrid[x+1][y] != -1 {
				repeatGrid[x+1][y] = -1
				step--
				backtrace(x+1, y, step, &repeatGrid)
				step++
				repeatGrid[x+1][y] = 0
			}

			// 向左
			if y-1 >= 0 && grid[x][y-1] != -1 && repeatGrid[x][y-1] != -1 {
				repeatGrid[x][y-1] = -1
				step--
				backtrace(x, y-1, step, &repeatGrid)
				step++
				repeatGrid[x][y-1] = 0
			}

			// 向右
			if y+1 < len(grid[x]) && grid[x][y+1] != -1 && repeatGrid[x][y+1] != -1 {
				repeatGrid[x][y+1] = -1
				step--
				backtrace(x, y+1, step, &repeatGrid)
				step++
				repeatGrid[x][y+1] = 0
			}

			return
		}
	}

	backtrace(start[0], start[1], needStep, &repeatGrid)

	return res
}
