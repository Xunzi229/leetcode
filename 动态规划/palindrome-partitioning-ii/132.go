/*
https://leetcode-cn.com/problems/palindrome-partitioning-ii/
*/
package palindrome_partitioning_ii

// 最少的分割次数 意味着最大的分割方式
// 从开始计算
func minCut(s string) int {
	if len(s)-1 <= 0 {
		return len(s) - 1
	}

	// 该DP存储的数据是切割的段数， 而非次数， 切割一次是 两段
	var dp = make([]int, len(s))
	dp[0] = 1

	// 使用动态规划
	// 每前进一步， 计算最短切割次数
	for i := 1; i <= len(s)-1; i++ {
		var min = 1<<31 - 1
		for j := 0; j <= i; j++ {
			if isPalindrome(s[j : i+1]) {
				if j == 0 {
					min = 1
					break
				} else {
					min = minValue(min, dp[j-1]+1)
				}
			}
		}
		dp[i] = min
	}

	return dp[len(s)-1] - 1
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	var start = 0
	var end = len(s) - 1
	for start <= end {
		if (len(s)%2 == 0) && (start+1 == end) && (s[start] == s[end]) {
			return true
		}

		if (len(s)%2 == 1) && (start == end) && (s[start] == s[end]) {
			return true
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return false
}
