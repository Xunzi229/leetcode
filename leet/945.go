//https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/
package main

import (
    "fmt"
    "sort"
)

func main(){
    a := []int{1, 2, 2}
    result := minIncrementForUnique(a)
    fmt.Println("move: ", result)
}

func minIncrementForUnique(A []int) int {
    sort.Ints(A)
    var numbers int
    for i, _ := range A {
        if i <= 0 {
            continue
        }
        for A[i] <= A[i-1] {
            A[i] = A[i] + 1
            numbers += 1
        }
    }
    return numbers
}
