/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
package longest_substring_without_repeating_characters

import "fmt"

// " "
// "au"
// "dvdf"
// "pwwkew"
// "abcabcbb"
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 如果将窗口内的函数标记位置
	wins := map[byte]int{}
	nums := map[byte]int{}

	// 记录一个最小值
	max := 1
	// left 是上一次出现重复字符的位置, right 是当前位置
	left, right := 0, 0

	// 遍历增加
	// 如果出现重复, 则该字符重置为1 且left当前的位置
	for right < len(s) {
		// 已经出现了
		nums[s[right]]++
		if nums[s[right]] > 1 {
			if wins[s[right]]+1 > left {
				left = wins[s[right]] + 1
			}
			nums[s[right]]--
		}
		wins[s[right]] = right
		if right-left+1 > max {
			fmt.Println(left, right)
			max = right - left + 1
		}
		right++
	}
	return max
}
