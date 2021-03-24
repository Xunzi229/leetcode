/*
https://leetcode-cn.com/problems/word-break/
*/
package word_break

/*
"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]


"leetcode"
["leet", "code"]

"ab"
["a","b"]

"applepenapple"
["apple","pen"]

"abcd"
["a","abc","b","cd"]
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictMap := isConclude(wordDict)

	dp := make([]bool, len(s))
	dp[0] = len(wordDictMap[s[0:1]]) > 0
	if len(s) == 1 {
		return dp[0]
	}

	for i := 1; i <= len(s)-1; i++ {
		var t = false
		for j := 0; j <= i; j++ {
			if j == 0 {
				if len(wordDictMap[s[j:i+1]]) > 0 {
					t = true
					break
				}
			} else {
				if dp[j-1] && len(wordDictMap[s[j:i+1]]) > 0 {
					t = true
					break
				}
			}
		}
		dp[i] = t
	}
	return dp[len(s)-1]
}

func isConclude(wordDict []string) map[string]string {
	ab := make(map[string]string)
	for _, v := range wordDict {
		ab[v] = v
	}
	return ab
}
