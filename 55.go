// https://leetcode-cn.com/problems/jump-game/

package main

import "fmt"

func main() {
    // 2,5,0,0
    // [1,1,2,2,0,1,1]
    // [3,2,1,0,4]
    
    fmt.Println(canJump([]int{3,2,1,0,4}))
}


//

/*
l := len(nums) - 1
for i := l-1; i >=0 ; i-- {
    if nums[i]+i >= l {
        l = i
    }
}
return l <= 0
*/

func canJump(nums []int) bool {
    length := len(nums)
    jumpNumbers := 0
    if length <= 1 || nums[0] + 1 >= length{
        return true
    }
    if nums[0] == 0 {
        return false
    }
    jumpNumbers += nums[0]
    for {
       nextStep := nums[jumpNumbers - 1]
       if nextStep == 0 {
           return false
       }
       fmt.Println(nextStep)
       if jumpNumbers + nextStep >= length {
           return true
       }
       jumpNumbers += nextStep + 1
    }
}