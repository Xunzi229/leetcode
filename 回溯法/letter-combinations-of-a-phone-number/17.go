//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

package letter_combinations_of_a_phone_number

/*
2：  abc
3:	def
4:	ghi
5:	jkl
6:  mno
7:  pqrs
8:  tuv
9:  wxyz
*/

/*
怎么思考这样的问题：
1. 首先是肯定不能重复的
2. 数字的个数代表了能 组合字母的长度
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	diet := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	total := make([]string, 0)

	var backtrace func(res string, digits string)

	backtrace = func(res string, digits string) {
		if 0 == len(digits) {
			total = append(total, res)
			return
		}

		for _, v := range diet[string(digits[0])] {
			res += string(v)
			backtrace(res, digits[1:])
			res = res[:len(res)-1]
		}
	}

	backtrace("", digits)

	return total
}
