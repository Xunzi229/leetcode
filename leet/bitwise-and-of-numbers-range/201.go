// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
package bitwise_and_of_numbers_range

// TODO
func rangeBitwiseAnd(m int, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
