/*
https://leetcode-cn.com/problems/minimum-window-substring/submissions/
*/
package minimum_window_substring

func minWindow(s string, t string) string {
    left, right := 0, 0
    windows := make([]string, 0)

    minStr := s
    for right < len(s) {
        windows = append(windows, string(s[right]))
        right++
        for valid(windows, t) && left < right {
            str := func() string {
                f := ""
                for _, v := range windows {
                    f += v
                }
                return f
            }()
            if len(str) < len(minStr) {
                minStr = str
            }

            windows = append([]string{}, windows[1:]...)
            left++
        }
    }
    return minStr
}

func valid(windows []string, t string) bool {
    for {
        for _, v := range t {
            vStr := string(v)

            c := false
            for i, v := range windows {
                if v == vStr {
                    c = true
                    if i == 0 {
                         windows = append([]string{}, windows[1:]...)
                    } else if i == len(windows) - 1 {
                        windows = append([]string{}, windows[0:(len(windows) - 1)]...)
                    } else {
                        windows = append(windows[0:i], windows[(i+1):]...)
                    }
                    break
                }
            }
            if !c {
                return false
            }
        }
        return true
    }
}