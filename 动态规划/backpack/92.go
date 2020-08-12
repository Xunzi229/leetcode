/*
https://www.lintcode.com/problem/backpack/description
*/
package backpack

/*
92. 背包问题
中文English
在n个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为m，每个物品的大小为A[i]

样例
样例 1:
	输入:  [3,4,8,5], backpack size=10
	输出:  9

样例 2:
	输入:  [2,3,5,7], backpack size=12
	输出:  12

挑战
O(n x m) 的时间复杂度 and O(m) 空间复杂度
如果不知道如何优化空间O(n x m) 的空间复杂度也可以通过.

注意事项
你不可以将物品进行切割。
*/

/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @return: The maximum size
 */
func backPack(m int, A []int) int {
	// 假设包的大小是变化的, 那么这个问题就是 就变成递归了

	dp := make([]int, m+1)
	dp[0] = 0 // 当包的大小为 0的时候是不能存放任何东西的

	for i := 1; i <= m; i++ {
		for j := 0; j <= len(A)-1; j++ {
			// 保证包不会超重
			if i-A[j] >= 0 {
				dp[i] = max(dp[i], dp[i-A[j]]+1)
			}
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
