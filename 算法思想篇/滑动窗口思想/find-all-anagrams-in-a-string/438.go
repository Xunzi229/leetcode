/*
https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
*/
package find_all_anagrams_in_a_string

/*
给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
*/

func findAnagrams(s string, p string) []int {
	wins := map[byte]int{} // 用于记录窗口数
	need := map[byte]int{}
	pLen := len(p)
	for i := 0; i < pLen; i++ {
		need[p[i]]++
	}

	sites := make([]int, 0) // 用于存储匹配到的位置

	left, right, match := 0, 0, 0
	for right < len(s) {
		x := s[right]
		right++
		if need[x] != 0 {
			wins[x]++
			if wins[x] == need[x] {
				match++
			}
		}

		for right-left >= pLen {
			// 重复字符, 这个时候匹配的长度其实是 map的长度
			if right-left == pLen && match == len(need) {
				sites = append(sites, left)
			}

			y := s[left]
			left++
			if need[y] != 0 {
				if wins[y] == need[y] {
					match--
				}
				wins[y]--
			}

		}
	}

	return sites
}
