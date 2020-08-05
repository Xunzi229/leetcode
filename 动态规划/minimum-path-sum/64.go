//https://leetcode-cn.com/problems/minimum-path-sum/
package minimum_path_sum

import "fmt"

func minPathSum(grid [][]int) int {
	//c := len(grid)
	//nums := make([][]int, c)
	//// 2、初始化
	//for i := 0; i < c; i++ {
	//	for j := 0; j < len(grid[i]); j++ {
	//		if nums[i] == nil {
	//			nums[i] = make([]int, len(grid[i]))
	//		}
	//		nums[i][j] = -1
	//	}
	//}
	//var result int
	_ok, result := minPathSumDeep(grid, 0, 0)
	fmt.Println(_ok)
	return result
}

// 右上角 到 右下角的位置
func minPathSumDeep(grid [][]int, x, y int) (bool, int) {
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
