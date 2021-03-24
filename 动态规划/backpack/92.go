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

// 如何思考, 思考的目的是什么,
// 最终想要的结果是 我能不能将当前的物品继续向包里面装
// 如何选择最优的解决方案
// 物品的个数是有限的(放完是不会放回去的), 包的重量是已知的

// 反思维: 包在装物品的同时, 也就是说这个物品的重量里面有这个包, 譬如 我物品重量是 5, 那么包是 5以后的都能装下这个包

// 保证新加的重量不能超重 , 其次 保证这个重量是合法的(如果当前重量是合法的, 那么上一次肯定也是合法的)
func backPack(m int, A []int) int {
	//dp[i][j] 表示i重量的时候能不能加上 j 重量
	dp := make([][]bool, len(A)+1) // + 1 长度上是 0, 0
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]bool, m+1) // + 1 长度上是 0, 0
	}
	dp[0][0] = true

	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]

			// j-A[i-1] >= 0 表示加上 j 会不会超重
			// dp[i-1][j-A[i-1]] // 表示这个是否是个可选的重量, 因为涉及是状态的变更, 我们不关心走到在哪一步, 我们只关心状态变更
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
		}
	}

	for i := m; i >= 0; i-- {
		if dp[len(A)][i] {
			return i
		}
	}
	return 0
}
