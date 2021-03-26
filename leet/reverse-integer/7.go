/*
https://leetcode-cn.com/problems/reverse-integer/
*/
package reverse_integer

import "math"

func reverse(x int) int {
	s := 1
	if x < 0 {
		s = -1
		x *= s
	}

	c := 0
	t := x % 10
	x = x / 10
	for t != 0 || x != 0 {
		if math.MaxInt32-t < c*10 {
			return 0
		}
		if s <= 0 {
			if math.MinInt32+t > c*10*s {
				return 0
			}
		}
		c = (c * 10) + t
		t = x % 10
		x = x / 10
	}
	return c * s
}
