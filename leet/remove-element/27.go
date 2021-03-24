// https://leetcode-cn.com/problems/remove-element/
package remove_element

/*
解题的意图是不再分配新的内存空间
然后返回数组的长度后这就是新的数组长度
*/
func removeElement(nums []int, val int) int {
    var x = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] ^ val == 0 {
            continue
        }
        nums[x] = nums[i]
        x++
    }
    return x
}