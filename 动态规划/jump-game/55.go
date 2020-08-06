/*
https://leetcode-cn.com/problems/jump-game/
*/
package jump_game

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。


还有一种可能中间有0 . 如果跳到这个点则不能到达终点
*/

// 动态规划
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	f := make([]bool, len(nums))
	f[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}

// 思维固化中的想法
//
// 能到最后一个点的点集
// 能到 点集的点集
// ...
// [2,3,1,1,4]
// [0,2,3]
// 这种方式超出了内存

func canJumpDeep(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	points := []int{len(nums) - 1}
	for {
		p := make([]int, 0)
		for _, pi := range points {
			for i := 0; i < pi; i++ {
				if nums[i]+i >= pi {
					// 能到达底部不需要再计算
					if i == 0 {
						return true
					}
					p = append(p, i)
				}
			}
		}
		if len(p) == 0 {
			return false
		}
		points = p
	}
}

// 贪心算法
// 假设能达到最后一个点
//
func canJumpDeep2(nums []int) bool {
	l := len(nums) - 1
	for i := l - 1; i >= 0; i-- {
		// 因为 i 在不断降低 而终点的位置其实可能一直没有变化. 如果有一个i 可以到达l点, 这个可能是最小的一个点
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}
