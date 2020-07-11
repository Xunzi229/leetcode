//https://leetcode-cn.com/problems/reverse-words-in-a-string/
package main

import (
    "fmt"
    "strings"
)

func main(){
    fmt.Println(reverseWords("  hello world!  "))
}


func reverseWords(s string) string {
    words := strings.Fields(s)

    length := len(words)
    for i := 0; i < length / 2; i++ {
        temp := words[length - 1 - i]
        words[length - 1 - i] = words[i]
        words[i] = string(temp)
    }

    return strings.Join(words, " ")
}