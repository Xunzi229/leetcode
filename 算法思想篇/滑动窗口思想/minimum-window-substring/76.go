/*
https://leetcode-cn.com/problems/minimum-window-substring/submissions/
*/
package minimum_window_substring

/*
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
*/

// 此种方式计算会超时
func minWindow(s string, t string) string {
	left, right := 0, 0
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
	minStr := s + t
	for right < len(s) {
		windows = append(windows, right)
		wMap[s[right]] += 1

		for validMap() && left <= right {
			str := ""
			for _, v := range windows {
				str += string(s[v])
			}

			if len(str) < len(minStr) {
				minStr = str
			}
			wMap[s[(windows[0])]] -= 1
			windows = windows[1:]
			left++
		}
		right++
	}
	if len(minStr) > len(s) {
		return ""
	}
	return minStr
}
