/*
https://leetcode-cn.com/problems/longest-common-subsequence/
*/
package longest_common_subsequence

/*
"abcde"
"ace"

"oxcpqrsvwf"
"shmtulqrypy"

"bl"
"xbl"

"bl"
"yby"

"ezupkr"
"ubmrapg"
*/
// 非常经典的二维的动态规划
// 使用矩阵可以发现规律
// TODO
func longestCommonSubsequence(a string, b string) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(a)][len(b)]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
