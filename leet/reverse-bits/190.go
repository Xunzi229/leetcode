//https://leetcode-cn.com/problems/reverse-bits/

package reverse_bits

// 颠倒二进制

// 涉及获取二进制最后一位
// 二进制的左移动  右移动
// 二进制的尾部追加 二进制的首部追加
func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	i := 32
	for {
		t := num & 1
		res = (res << 1) + t
		num = num >> 1

		i--
		if num == 0 {
			res = res << i
			break
		}
	}
	return res
}
