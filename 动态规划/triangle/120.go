//https://leetcode-cn.com/problems/triangle/
package triangle

// 动态规划自底向上
func minimumTotal(triangle [][]int) int {
	/*
		// DFS 方式
		return minimumTotalDFS(triangle)
	*/

	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1 定一个一个函数表示每一层的位置 function[i][j]
	// 初始化
	c := len(triangle)
	function := make([][]int, c)
	// 2、初始化
	for i := 0; i < c; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if function[i] == nil {
				function[i] = make([]int, len(triangle[i]))
			}
			function[i][j] = triangle[i][j]
		}
	}

	// 自底层迭代
	// 顶层不用计算
	for i := c - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			function[i][j] = min(function[i+1][j], function[i+1][j+1]) + triangle[i][j]
		}
	}

	// 获取顶层的最短路径
	return function[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
// 使用 DFS, 记忆话搜索, 这里已经是用到了动态规划
func minimumTotalDFS(triangle [][]int) int {
	c := len(triangle)
	nums := make([][]int, c)
	// 2、初始化
	for i := 0; i < c; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if nums[i] == nil {
				nums[i] = make([]int, len(triangle[i]))
			}
			nums[i][j] = -1
		}
	}
	return minimumTotalDFSDeep(nums, triangle, 0, 0)
}

func minimumTotalDFSDeep(nums, triangle [][]int, x, y int) int {
	// 缓存结果
	if nums[x][y] != -1 {
		return nums[x][y]
	}
	if len(triangle)-1 == x {
		return triangle[x][y]
	}

	//
	l := minimumTotalDFSDeep(nums, triangle, x+1, y)
	r := minimumTotalDFSDeep(nums, triangle, x+1, y+1)

	nums[x][y] = min(l, r) + triangle[x][y]
	return nums[x][y]
}
*/

// 自上而下
func minimumTotalW(triangle [][]int) int {
	c := len(triangle)
	nums := make([][]int, c)
	// 2、初始化
	for i := 0; i < c; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if nums[i] == nil {
				nums[i] = make([]int, len(triangle[i]))
			}
			nums[i][j] = triangle[i][j]
		}
	}

	for i := 1; i < c; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 < 0 {
				nums[i][j] = nums[i-1][j] + triangle[i][j]
			} else if j >= len(nums[i-1]) {
				nums[i][j] = nums[i-1][j-1] + triangle[i][j]
			} else {
				nums[i][j] = min(nums[i-1][j], nums[i-1][j-1]) + triangle[i][j]
			}
		}
	}

	result := nums[c-1][0]
	for i := 1; i < len(nums[c-1]); i++ {
		result = min(result, nums[c-1][i])
	}
	return result
}
