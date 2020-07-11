//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/submissions/
// 解题思路， 首先肯定是需要遍历这个字符串的
// 我们可已想想一个框 框住最长的一组字符串， 框需要左边的坐标位置 和 右边的坐标位置
// 假设我们中间移动一个位置
// 我们需要判断这个位置的值是否在前面的框中 ，框我们可以用一组hash 来存 （字符当做key值， 其他的信心如 位置和数量用一个结构体存起来）， 因为hash的key不能重复， 所以不用担心不唯一的问题
// 检查key是否存在， 如果存在的话，且这个key需要在框中，  需要将 框的左位置 加1， 右位置当前的坐标， 判断长度记录一个最大框时候的位置
// 如果key不存在， 则继续存这个下去， 更新或者保存这个值得位置和，并且要计算最后的长度，如果大于记录的最大框的长度，则需要更新
// 如此下来就能在O(n)的时间复杂度下完成这个任务

package main

import (
    "fmt"
)
func main(){
    s :=  "abcdeeeabcdeefgff"
    fmt.Println(lengthOfLongestSubstring(s))
}
type Letter struct {
    LastSite int
    Number   int
}
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0{
        return 0
    }
    var MaxLen int
    var left, right = 0, 0

    MaxLen = 1
    var mapDir = make(map[string]*Letter)
    for i, v := range s {
        if mapDir[string(v)] != nil && mapDir[string(v)].LastSite >= left {
            if right - left + 1 > MaxLen {
                MaxLen = right - left + 1
            }
            left = mapDir[string(v)].LastSite + 1
            right = i
            mapDir[string(v)].LastSite = i
            continue
        }
        if mapDir[string(v)] != nil {
            mapDir[string(v)].LastSite = i
        } else {
           mapDir[string(v)] = &Letter{
                Number: 1,
                LastSite: i,
           }
        }
        right = i
        if right - left + 1 > MaxLen {
            MaxLen = right - left + 1
        }
    }
    return MaxLen
}