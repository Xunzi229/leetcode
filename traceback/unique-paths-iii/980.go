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
	x0, y0 := 0, 0
	x1, y1 := 0, 0
	needStep := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				x0 = x
				y0 = y
				grid[x][y] = -1
			}
			if grid[x][y] == 2 {
				x1 = x
				y1 = y
			}
			if grid[x][y] == 0 {
				needStep++
			}
		}
	}

	var res int

	var backtrace func(x, y, step int)

	backtrace = func(x, y, step int) {
		// 到达终点的标志, 为啥等于 -1, 因为 0 的表示都走完了，最后一步到这个地方就变成-1了
		if x == x1 && y == y1 {
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
			if x-1 >= 0 && grid[x-1][y] != -1 {
				if !(step != 0 && grid[x-1][y] == 2) {
					grid[x-1][y] = -1
					step--
					backtrace(x-1, y, step)
					step++
					grid[x-1][y] = 0
				}
			}

			// 向下
			if x+1 < len(grid) && grid[x+1][y] != -1 {
				if !(step != 0 && grid[x+1][y] == 2) {
					grid[x+1][y] = -1
					step--
					backtrace(x+1, y, step)
					step++
					grid[x+1][y] = 0
				}
			}

			// 向左
			if y-1 >= 0 && grid[x][y-1] != -1 {
				if !(step != 0 && grid[x][y-1] == 2) {
					grid[x][y-1] = -1
					step--
					backtrace(x, y-1, step)
					step++
					grid[x][y-1] = 0
				}
			}

			// 向右
			if y+1 < len(grid[x]) && grid[x][y+1] != -1 {
				if !(step != 0 && grid[x][y+1] == 2) {
					grid[x][y+1] = -1
					step--
					backtrace(x, y+1, step)
					step++
					grid[x][y+1] = 0
				}
			}
			return
		}
	}

	backtrace(x0, y0, needStep)

	return res
}
