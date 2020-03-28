//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/submissions/
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