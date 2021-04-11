/*
https://leetcode-cn.com/problems/word-search/
*/
package word_search

func exist(board [][]byte, word string) bool {
	var result = false

	var backtrace func(n, x, y int, w string)
	colorBoard := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		colorBoard[i] = make([]int, len(board[i]))
	}

	backtrace = func(n, x, y int, w string) {
		if w == word {
			result = true
			return
		}

		if x < 0 || y < 0 || x > len(board) || y > len(board[x]) {
			return
		}

		var nextStr byte
		if n+1 < len(word) {
			nextStr = word[n+1]
		}

		// 寻找下一个位置
		// 保证下一个位置可以触达， 不能超出边界
		// 保证下个位置没有重复
		// 保证下个位置是单词内的字符
		if x-1 >= 0 && colorBoard[x-1][y] == 0 &&
			board[x-1][y] == nextStr {
			w += string(board[x-1][y])
			colorBoard[x-1][y] = 1
			backtrace(n+1, x-1, y, w)
			if result {
				return
			}
			colorBoard[x-1][y] = 0
			w = w[:len(w)-1]
		}

		if y-1 >= 0 && colorBoard[x][y-1] == 0 &&
			board[x][y-1] == nextStr {
			w += string(board[x][y-1])
			colorBoard[x][y-1] = 1
			backtrace(n+1, x, y-1, w)
			if result {
				return
			}
			colorBoard[x][y-1] = 0
			w = w[:len(w)-1]
		}

		if x+1 < len(board) && colorBoard[x+1][y] == 0 &&
			board[x+1][y] == nextStr {
			w += string(board[x+1][y])
			colorBoard[x+1][y] = 1
			backtrace(n+1, x+1, y, w)
			if result {
				return
			}
			colorBoard[x+1][y] = 0
			w = w[:len(w)-1]
		}

		if y+1 < len(board[x]) && colorBoard[x][y+1] == 0 &&
			board[x][y+1] == nextStr {
			w += string(board[x][y+1])
			colorBoard[x][y+1] = 1
			backtrace(n+1, x, y+1, w)
			if result {
				return
			}
			colorBoard[x][y+1] = 0
			w = w[:len(w)-1]
		}
		return
	}

	// 起始位置判断
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if result {
				return result
			}
			if board[x][y] == word[0] {
				colorBoard[x][y] = 1
				backtrace(0, x, y, string(word[0]))
				colorBoard[x][y] = 0
			}
		}
	}

	return result
}
