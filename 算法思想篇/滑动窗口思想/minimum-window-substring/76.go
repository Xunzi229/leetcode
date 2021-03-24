/*
https://leetcode-cn.com/problems/minimum-window-substring/submissions/
*/
package minimum_window_substring

// 此种方式计算会超时
func minWindow(s string, t string) string {
	windows := make([]int, 0)

	tMap := map[byte]int{}
	for i, _ := range t {
		tMap[t[i]] += 1
	}
	wMap := map[byte]int{}

	validMap := func() bool {
		for kt, vt := range tMap {
			if wMap[kt] < vt {
				return false
			}
		}
		return true
	}

	minLen := 1<<31 - 1
	r, l := -1, -1
	for left, right := 0, 0; right < len(s); right++ {
		windows = append(windows, right)
		wMap[s[right]] += 1

		for validMap() && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				r, l = right+1, left
			}
			wMap[s[(windows[0])]] -= 1
			windows = windows[1:]
			left++
		}
	}
	if l == -1 {
		return ""
	}
	return s[l:r]
}
