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
	watch := make(map[string]string)

	var backtrace func(ps, num int)

	// 前四位表示的上面的4个 LED
	// 后6个表示的是底部的 LED
	chess := make([]int, 10) // 表示的是当前的一个状态
	// num 表示剩余, c 表示当前的位置
	backtrace = func(c, n int) {
		if n == 0 || c == 10 { // 当前位置超过了， 或者 已经没有可以
			if v, ok := calc(chess); ok {
				watch[v] = v
			}
			return
		}

		for n != 0 && c < 10 {
			chess[c] = 1
			backtrace(c+1, n-1)
			chess[c] = 0
			c++
		}
	}

	for i := 0; i <= 0; i++ {
		if i+num <= 9 {
			backtrace(i, num)
		}
	}
	r := make([]string, 0)
	for _, v := range watch {
		r = append(r, v)
	}
	return r
}

// 判断最终值
func calc(chess []int) (r string, ok bool) {
	h := 0
	m := 0
	for i := 0; i < len(chess); i++ {
		if i < 4 {
			h += int(math.Pow(2, float64(i))) * chess[i]
			if h > 11 {
				return "", false
			}
		} else {
			m += int(math.Pow(2, float64(i-4))) * chess[i]
			if m > 59 {
				return "", false
			}
		}

	}
	return fmt.Sprintf("%d:%02d", h, m), true
}
