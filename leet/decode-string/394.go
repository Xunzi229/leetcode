//https://leetcode-cn.com/problems/decode-string/

package decode_string

import (
	"strconv"
	"strings"
)

// "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
// "3[a]2[bc]"
// "3[a2[c]]"
//  "2[abc]3[cd]ef"
// "abc3[cd]xyz"
func decodeString(s string) string {
	// 使用栈来存储值
	stack := make([]string, 0)
	for _, _v := range s {
		v := string(_v)
		// 暂时不考虑奇怪的字符, 这是一个正常的压缩的字符串

		// 先考虑出栈的情况
		if v == "]" {
			in := make([]string, 0)
			for {
				val := pop(&stack)
				if val == "[" {
					break
				}
				in = append([]string{val}, in...)
			}
			str := strings.Join(in, "")

			numStr := ""
			for {
				if len(stack) == 0 {
					break
				}
				val := pop(&stack)

				valByte := []byte(val)
				if valByte[0] > 58 || valByte[0] < 48 {
					stack = append(stack, val)
					break
				}
				numStr = val + numStr
			}
			num, _ := strconv.Atoi(numStr)
			res := strAdd(str, num)
			stack = append(stack, res)
			continue
		}
		stack = append(stack, v)
	}

	return strings.Join(stack, "")
}

func strAdd(str string, num int) string {
	emStr := ""
	for i := num; i >= 1; i-- {
		emStr += str
	}
	return emStr
}

func pop(stack *[]string) string {
	if len(*stack) == 0 {
		return ""
	}

	v := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return v
}
