// https://leetcode-cn.com/problems/longest-palindromic-substring/

package main

import "fmt"

func main(){
    fmt.Println(longestPalindrome("abacddcabaef"))
}

func longestPalindrome(s string) string {
    sLen := len(s)
    if sLen <= 1 {
        return s
    }
    for x := sLen; x >= 0; x-- {
        for startPoint := 0; startPoint <= sLen - x - 1; startPoint++ {
            i := startPoint

            j := x + i -1
            var can bool = true
            for can {
                if i > sLen -1 || j <= 0 {
                    can = false
                    break
                }
                if s[i] != s[j] {
                   can = false
                   break
                }
                if i == j &&  x % 2 == 1 {
                    break
                }
                if i + 1 == j &&  x % 2 == 0{
                    break
                }
                i++
                j--
            }
            if can {
                return s[startPoint:(startPoint + x)]
            }

        }
    }
    return ""
}