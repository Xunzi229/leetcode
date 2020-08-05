// https://leetcode-cn.com/problems/climbing-stairs/
package climbing_stairs

/*
当n等于1的时候，只需要跳一次即可，只有一种跳法，记f(1)=1
当n等于2的时候，可以先跳一级再跳一级，或者直接跳二级，共有2种跳法，记f(2)=2
当n等于3的时候，他可以从一级台阶上跳两步上来，也可以从二级台阶上跳一步上来，所以总共有f(3)=f(2)+f(1)；
等于n的时候，总共有f(n)=f(n-1)+f(n-2)(这里n>2)种跳法。
*/
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[n] = f[n-1] + f[n-2]
	}
	return f[n]
}
