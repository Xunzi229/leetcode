/*
https://leetcode-cn.com/problems/permutation-in-string/
*/
package permutation_in_string

/*
给定两个字符串s1和s2，写一个函数来判断 s2 是否包含 s1的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False
*/
// 此种方式已经超时了. 有没有更好的解决方式
// 如果遇到其中一个字符串不在S1中, 需要跳过这个字符串, 以该字符串为起点是不是?
//
/*
func checkInclusion(s1 string, s2 string) bool {
    s1A := strings.Split(s1, "")
    sort.Strings(s1A)
    s1 = strings.Join(s1A, "")

    for left, right := 0, len(s1); right <= len(s2); right++ {
        w := s2[left:right]
        fmt.Println(w)
        if validateCompare(s1, w) {
            return true
        }
        left++
    }
    return false
}

func validateCompare(s1, w string) bool {
	ws := strings.Split(w, "")
	sort.Strings(ws)
	return strings.Join(ws, "") == s1
}

*/

func checkInclusion(s1 string, s2 string) bool {
	ws := make(map[byte]int) // 窗口, 用于备份窗口内的数量,
	ns := make(map[byte]int) // 需要的,

	for i := 0; i < len(s1); i++ {
		ns[s1[i]]++
	}

	left, right := 0, 0
	match := 0 // 用与标记窗口内字符串相符的个数, 如果相符个数等于s1 长度说明全匹配

	for right < len(s2) {
		c := s2[right]
		right++
		if ns[c] != 0 {
			ws[c]++
			if ws[c] == ns[c] {
				match++
			}
		}

		for right-left >= len(s1) {
			if match == len(ns) {
				return true
			}

			d := s2[left]
			left++
			if ns[d] != 0 {
				if ws[d] == ns[d] {
					match--
				}
				ws[d]--
			}
		}
	}
	return false
}
