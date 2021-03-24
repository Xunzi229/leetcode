/*
https://leetcode-cn.com/problems/decode-ways/
*/
package decode_ways

// 0 => 48
// TODO
// 状态转移  dp[i] = dp[i-1] + dp[i-2]
// 考虑数值 0 在其中的 作用
// 如果 dp[i-1] 和 dp[i-2] 都是为
// 如果 dp[i-1] = 0 , 则 dp[i-2]
// 如果 dp[i-2] = 0 则 dp[i-1]

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0
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

// 少存储点空间, 只使用前后变量变化
func numDecodingsWithOutSpace(s string) int {
	var cur, preStep1, preStep2 = 0, 0, 0
	cur = 0
	if s[0] != 48 {
		cur = 1
	}

	if len(s) <= 1 {
		return cur
	}
	preStep1, preStep2 = cur, cur

	for i := 2; i <= len(s); i++ {
		if s[i-2] == 48 && s[i-1] == 48 {
			return 0
		}
		num := (s[i-2]-48)*10 + (s[i-1] - 48)
		// "1011"
		if s[i-2] == 48 || num > 26 {
			cur = preStep1
		} else if s[i-1] == 48 {
			// "501"
			if num > 26 {
				return 0
			}
			cur = preStep2
		} else {
			// "111"
			cur = preStep1 + preStep2
		}
		preStep1, preStep2 = cur, preStep1
	}
	return cur
}
