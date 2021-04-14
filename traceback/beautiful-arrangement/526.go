/*
https://leetcode-cn.com/problems/beautiful-arrangement/
*/

package beautiful_arrangement

func countArrangement(n int) int {
	if n < 1 || n > 15 {
		return 0
	}

	res := 0

	var backtrace func(c, s int)
	var mapExist = make(map[int]int)
	// c 表示当前的位置(index+1)
	// s 表示当前位置的值
	backtrace = func(index, s int) {
		if index != 0 && s != 0 {
			if index == n {
				res++
				return
			}
		}
		for i := 1; i <= n; i++ {
			if mapExist[i] == 0 && ((index+1)%i == 0 || i%(index+1) == 0) {
				index++
				mapExist[i]++
				backtrace(index, i)
				index--
				mapExist[i]--
			}
		}
	}

	backtrace(0, 1)

	return res
}

func isExist(v int, num []int) bool {
	for _, nu := range num {
		if nu == v {
			return true
		}
	}
	return false
}
