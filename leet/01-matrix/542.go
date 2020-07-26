// https://leetcode-cn.com/problems/01-matrix/
package _1_matrix

func updateMatrix(matrix [][]int) [][]int {
	stacks := make([][]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				stacks = append(stacks, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(stacks) != 0 {
		point := stacks[0]
		stacks = stacks[1:]

		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				stacks = append(stacks, []int{x, y})
			}
		}
	}
	return matrix
}
