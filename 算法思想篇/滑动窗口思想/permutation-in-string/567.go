/*
https://leetcode-cn.com/problems/permutation-in-string/
*/
package permutation_in_string

import (
    "sort"
    "strings"
)

/*
给定两个字符串s1和s2，写一个函数来判断 s2 是否包含 s1的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False
*/
// 此种方式已经超时了. 有没有更好的解决方式
// 如果遇到其中一个字符串不在S1中, 需要跳过这个字符串, 以该字符串为起点是不是?
//
/*
func checkInclusion(s1 string, s2 string) bool {
    s1A := strings.Split(s1, "")
    sort.Strings(s1A)
    s1 = strings.Join(s1A, "")

    for left, right := 0, len(s1); right <= len(s2); right++ {
        w := s2[left:right]
        fmt.Println(w)
        if validateCompare(s1, w) {
            return true
        }
        left++
    }
    return false
}
*/

// 跳过不需要的字段
func checkInclusion(s1 string, s2 string) bool {
    s1A := strings.Split(s1, "")
    sort.Strings(s1A)
    s1 = strings.Join(s1A, "")

    mapS1 := map[byte]int{}
    for i, _ := range s1 {
		mapS1[s1[i]] += 1
	}

    left, right := 0, len(s1)
    for right <= len(s2){
        if mapS1[s2[right-1]] == 0 {
            left = right
            right += len(s1)
            continue
        }
        w := s2[left:right]
        if validateCompare(s1, w) {
            return true
        }
        left++
        right++
    }
    return false
}

func validateCompare(s1, w string) bool {
    ws := strings.Split(w, "")
    sort.Strings(ws)
    return strings.Join(ws, "") == s1
}