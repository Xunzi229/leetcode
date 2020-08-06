/*
https://leetcode-cn.com/problems/jump-game-ii/
*/

package jump_game_ii

// 贪心算法-.
// 有一个特性是 如果A 能到达C 那么 A 肯定能到达 B, 所以找第一个能到达C的点
// 如果能找A点, 则再找能最先到达A点的点
//
func jump(nums []int) int {
	l := len(nums) - 1
	if l <= 0 {
		return 0
	}

	jump := 0
	for l >= 0 {
		for i := 0; i < l; i++ {
			if i == 0 && nums[i]+i >= l {
				return jump + 1
			}
			if nums[i]+i >= l {
				l = i
				jump++
				break
			}
			if i == l-1 {
				return 0
			}
		}
	}
	return 0
}
