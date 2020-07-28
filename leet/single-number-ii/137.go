//https://leetcode-cn.com/problems/single-number-ii/

package single_number_ii

// 取巧的计算方式将数转化成二进制的数, 统计相同位的数的值 取余
/*
	   1 1 1 1  => 15
       1 1 0 1  => 13
       1 1 1 1  => 15
       1 1 1 1  => 15
---------------
	   4 4 3 4   /  3
--------------------
       1 1 0 1

*/
func singleNumber(nums []int) int {
	var result int
	for i := 63; i >= 0; i-- {
		var sum = 0
		for _, v := range nums {
			sum += (v >> i) & 1
		}
		result = (result << 1) | (sum % 3)
	}
	return result
}
