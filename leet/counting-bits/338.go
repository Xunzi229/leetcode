//https://leetcode-cn.com/problems/counting-bits/

package counting_bits

func countBits(num int) []int {
	n := make([]int, num+1)

	for i := 0; i <= num; i++ {
		n[i] = count(i)
	}
	return n
}

func count(num int) int {
	var res int
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}

// 动态规划的解法
// TODO
