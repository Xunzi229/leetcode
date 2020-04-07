//https://leetcode-cn.com/problems/rotate-matrix-lcci/
//面试题 01.07. 旋转矩阵

// 好像很容易 ， 只要旋转两次就行， 首先以对角翻转，
// 再以横向中轴翻转


package main

import "fmt"

func main() {
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    rotate(matrix)

    fmt.Println(matrix)
}

func rotate(matrix [][]int)  {
    n := len(matrix)

    // 对折翻转
    i := 0; j := 0
    for i <= n - 1 && j <= n -1 {
        gi, gj := i - 1, j - 1
        for gi >= 0 && gj >= 0 {
           t := matrix[i][gj]
           matrix[i][gj] = matrix[gi][j]
           matrix[gi][j] = t
           gi--; gj--
        }
        i++; j++
    }


    // 横向翻转
    for j := 0 ; j < n; j++ {
       i := 0
       for true{
           if i >= n - i - 1 {
               break
           }
           t := matrix[j][i]
           matrix[j][i] = matrix[j][n - i - 1]
           matrix[j][n - i - 1] = t
           i++
       }
    }
}


