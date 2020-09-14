/*
https://leetcode-cn.com/problems/regular-expression-matching/
*/

package regular_expression_matching

/*
   保证S 可被 P 匹配到
   '.' 匹配任意单个字符 46
   '*' 匹配零个或多个前面的那一个元素 42
*/

/*
. 表示的是 匹配  一个  字符
* 表示匹配 0 or n 个 前面的字符   a* 可以匹配   "", "aa", "aaa", "aaa" (这个地方0个很骚)

用 i, j 标识 S P 字符串的位置

如果  j 位置是: a-z, 则这个地方 如果  f[i][j] = s[i] == p[j]
如果  j 位置是: *    则             f[i][j] = f[i][j-2]

\
如果开始
如何状态转移
如何结束

两个字符串相匹配,
	x 是 s 的某个位置
	y 是 p 的某个位置

x,y 如果是匹配的, 那么 x-1,  y-1 也应该需要匹配
...
这样才能保证 f[x][y] 是可以 匹配的

在考虑的, 因为 P 匹配到 S , 可以理解 P 包含了 S, (我怎么想到了 KMP)

怎么确定 f[x][y] 地方的值


          | f[i-1][j] or f[i][j-2],  s[i]=p[j−1]
f[i][j] = |
          | f[i][j-2], s[i] != p[j−1]

// 如何考虑这个问题
画一个EXCEL 可以看到状态转移的过程
s = "aa"
p = "a*"
*/

func isMatch(s string, p string) bool {
	// 增加边界初始值
	p = " " + p
	s = " " + s

	dp := make([][]bool, len(s))
	for i := 0; i <= len(s)-1; i++ {
		dp[i] = make([]bool, len(p))
	}
	dp[0][0] = true

	// 初始化边界值, 不考虑 最开始就是*字符这种无效字符情况
	for j := 1; j <= len(p)-1; j++ {
		if p[j] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 状态转移的过程
	for i := 1; i <= len(s)-1; i++ {
		for j := 1; j <= len(p)-1; j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if p[j] == '*' {
					dp[i][j] = dp[i][j-2]
					if !dp[i][j] {
						if dp[i-1][j] && (p[j-1] == s[i] || p[j-1] == '.') {
							dp[i][j] = true
						}
					}
				}
			}
		}
	}
	return dp[len(s)-1][len(p)-1]
}
