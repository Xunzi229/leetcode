// https://leetcode-cn.com/problems/number-of-1-bits/submissions/
// https://baike.baidu.com/item/%E6%B1%89%E6%98%8E%E9%87%8D%E9%87%8F
package number_of_1_bits

/*
1010101010101010101010101010101  0x55555555
0110011001100110011001100110011  0x33333333
0001111000011110000111100001111  0x0F0F0F0F
0000001000000010000000100000001  0x01010101
*/
func hammingWeight(num uint32) int {
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num & 0x0F0F0F0F) + ((num >> 4) & 0x0F0F0F0F)
	num = num * (0x01010101) >> 24
	return int(num)
}

// 这种方式可以获取最后一个1
func hammingWeightStyle2(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}
