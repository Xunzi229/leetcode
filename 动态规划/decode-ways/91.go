/*
https://leetcode-cn.com/problems/decode-ways/
*/
package decode_ways

// 0 => 48
// TODO
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != 48 {
		dp[1] = 1
	}
	if len(s) <= 1 {
		return dp[1]
	}

	for i := 2; i <= len(s); i++ {
		num := (s[i-2]-48)*10 + (s[i-1] - 48)
		if s[i-2] == 48 && s[i-1] == 48 {
			return 0
		} else if s[i-2] == 48 {
			dp[i] = dp[i-1]
		} else if s[i-1] == 48 {
			if num > 26 {
				return 0
			}
			dp[i] = dp[i-2]
		} else if num > 26 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[len(dp)-1]
}
