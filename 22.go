//https://leetcode-cn.com/problems/generate-parentheses/

package main

import "fmt"

func main (){
    fmt.Println(generateParenthesis(3))
}

// 此题使用回溯算法很好解
func generateParenthesis(n int) []string {
    result := make([]string, 0)
    backtrack(n, n, "",  &result)
    return result
}

func backtrack(left, right int, tmp string, result *[]string) {
    if right == 0 {
        *result = append(*result, tmp)
        return
    }

    if left > 0 {
        backtrack(left - 1, right, tmp + "(", result)
    }

    if right > left {
        backtrack(left, right - 1, tmp + ")", result)
    }
}