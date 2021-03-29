/*
https://leetcode-cn.com/problems/binary-watch/
*/
package binary_watch

import (
	"fmt"
	"math"
)

/*
401. 二进制手表
*/
// TODO
func readBinaryWatch(num int) []string {
	watch := make([]string, 0)

	var backtrace func(ps, num int)

	// 前四位表示的上面的4个 LED
	// 后6个表示的是底部的 LED
	chess := make([]int, 10)
	// num 表示剩余, c 表示当前的位置
	backtrace = func(c, n int) {
		if n == 0 || c == 10 {
			if v, ok := calc(chess); ok {
				watch = append(watch, v)
			}
		}
		for i := c; i < c+n; i++ {
			chess[i] = 1
			backtrace(i+1, n-1)
			chess[i] = 0
		}
	}
	for i := 0; i < len(chess); i++ {
		backtrace(i, num)
	}

	return watch
}

// 判断最终值
func calc(chess []int) (r string, ok bool) {
	if len(chess) != 10 {
		return "", false
	}
	h := 0
	m := 0
	for i := 0; i < len(chess); i++ {
		if i < 4 {
			h = h + int(math.Pow(2, float64(i)))*chess[i]
			if h > 11 {
				return "", false
			}
		} else {
			m = m + int(math.Pow(2, float64(i-4)))*chess[i]
			if h > 59 {
				return "", false
			}
		}

	}
	return fmt.Sprintf("%d:%02d", h, m), true
}
