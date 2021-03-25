/*
https://leetcode-cn.com/problems/roman-to-integer/submissions/
*/

package roman_to_integer

func romanToInt(s string) int {
	ref := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 1 {
		return ref[[]byte(s)[0]]
	}

	total := 0
	for i, _ := range s {
		if len(s) <= i+1 {
			total += ref[s[i]]
			return total
		}
		if ref[s[i]] < ref[s[i+1]] {
			total -= ref[s[i]]
			continue
		}
		total += ref[s[i]]
	}
	return total
}
