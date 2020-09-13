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



边界值 f[0][0]=true
*/
func isMatch(s string, p string) bool {
	return false
}
