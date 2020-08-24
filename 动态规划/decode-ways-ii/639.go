/*
https://leetcode-cn.com/problems/decode-ways-ii/
*/
package decode_ways_ii

/*
一条包含字母 A-Z 的消息通过以下的方式进行了编码：
'A' -> 1
'B' -> 2
...
'Z' -> 26
除了上述的条件以外，现在加密字符串可以包含字符 '*'了，字符'*'可以被当做1到9当中的任意一个数字。
给定一条包含数字和字符'*'的加密信息，请确定解码方法的总数。
同时，由于结果值可能会相当的大，所以你应当对109 + 7取模。（翻译者标注：此处取模主要是为了防止溢出）

示例 1 :
输入: "*"
输出: 9
解释: 加密的信息可以被解密为: "A", "B", "C", "D", "E", "F", "G", "H", "I".

示例 2 :
输入: "1*"
输出: 9 + 9 = 18（翻译者标注：这里1*可以分解为1,* 或者当做1*来处理，所以结果是9+9=18）
说明 :

输入的字符串长度范围是 [1, 105]。
输入的字符串只会包含字符 '*' 和 数字'0' - '9'。
*/

//
//
// "*" => 42
// "0" => 48
// "1" => 49
func numDecodings(s string) int {
	n := len(s)
	mod := int(1e9) + 7
	if n == 0 || s[0] == 48 {
		return 0
	}
	dp := make([]int, len(s)+1)
	a := " " + s
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if a[i] == 42 {
			dp[i] = (9 * dp[i-1]) % mod
			if a[i-1] == 49 {
				dp[i] = dp[i] + 9*dp[i-2]%mod
			} else if a[i-1] == 50 {
				dp[i] = dp[i] + 6*dp[i-2]%mod
			} else if a[i-1] == 42 {
				dp[i] = dp[i] + 15*dp[i-2]%mod
			}
		} else if a[i] == 48 {
			if a[i-1] == 49 {
				dp[i] = dp[i-2]
			} else if a[i-1] == 50 {
				dp[i] = dp[i-2]
			} else if a[i-1] == 42 {
				dp[i] = 2 * dp[i-2] % mod
			} else {
				return 0
			}
		} else {
			dp[i] = dp[i-1]
			if a[i-1] == 49 {
				dp[i] = (dp[i] + dp[i-2]) % mod
			} else if a[i-1] == 50 && a[i]-48 >= 1 && a[i]-48 <= 6 {
				dp[i] = (dp[i] + dp[i-2]) % mod
			} else if a[i-1] == 42 {
				if a[i]-42 >= 7 && a[i]-48 <= 9 {
					dp[i] = (dp[i] + dp[i-2]) % mod
				} else {
					dp[i] = (dp[i] + 2*dp[i-2]) % mod
				}
				dp[i] = 2 * dp[i-2] % mod
			}
		}
	}
	return dp[n] % mod
}
