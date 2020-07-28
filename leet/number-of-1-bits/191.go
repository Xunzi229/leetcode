// https://leetcode-cn.com/problems/number-of-1-bits/submissions/
package number_of_1_bits

func hammingWeight(num uint32) int {
	var c = 0
	for i := 0; i < 64; i++ {
		if (num >> i & 1) == 1 {
			c++
		}
	}
	return c
}
