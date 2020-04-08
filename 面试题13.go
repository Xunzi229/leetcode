//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
package main

import (
    "fmt"
    "strconv"
)

func main(){
    fmt.Println(movingCount(4, 5, 0))
}

// 使用一个二维数组， 初始化数组的值都是 0 ， 如果能到达则标记 1 ，如果不能到达， 还是为 0
func movingCount(m int, n int, k int) int {
    var chess = make([][]int, m)

    for  i := 0; i < m; i++ {
          for j := 0; j < n; j++ {
              chess[i] = append(chess[i] , 0)
          }
    }
    canTotal := 0
    if k == 0 {
        return 1
    }
    changeState1(0, 0, k, m - 1, n - 1, &canTotal, chess)

    return canTotal
}

//x y 是当前的位置
func changeState1(x, y, k, totalX, totalY int, canTotal *int, chess [][]int){
    // 左
    if x >= 0 && y >= 0 && x - 1 >= 0 && y <= totalY {
        if chess[x-1][y] != 1 && chess[x-1][y] != 3{
            chess[x-1][y] = 1
            if sumParseInt(x-1)+sumParseInt(y) <= k {
                *canTotal += 1
                changeState1(x-1, y, k, totalX, totalY, canTotal, chess)
            }
        }
    }
    // 右
    if x >= 0 && y >= 0 && x + 1 <= totalX && y <= totalY {
        if chess[x + 1][y] != 1 && chess[x + 1][y] != 3{
            chess[x + 1][y] = 1
            if sumParseInt(x + 1) + sumParseInt(y) <= k {
                *canTotal += 1
                changeState1(x + 1, y, k, totalX, totalY, canTotal, chess)
            }
        }

    }
    // 上
    if x >= 0 && y >= 0 && x <= totalX && y + 1 <= totalY {
        if chess[x][y + 1] != 1{
            chess[x][y + 1] = 1
            if sumParseInt(x) + sumParseInt(y + 1) <= k {
                *canTotal += 1
                changeState1(x, y + 1, k, totalX, totalY, canTotal, chess)
            }
        }

    }
    // 下
    if x >= 0 && y >= 0 && x <= totalX && y - 1 >= 0 {
        if chess[x][y - 1] != 1{
           chess[x][y - 1] = 1
           if sumParseInt(x) + sumParseInt(y - 1) <= k {
                *canTotal += 1
                changeState1(x, y - 1, k, totalX, totalY, canTotal, chess)
           }
        }
    }
}

func sumParseInt(x int) int{
    str := strconv.Itoa(x)
    total := 0
    for _ , v := range str {
        if i, err := strconv.Atoi(string(v)); err == nil {
            total += i
        }
    }
    return total
}