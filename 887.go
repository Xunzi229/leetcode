package main

import "fmt"

// 动态规划
// 题目的意思 是 如果一定要测出F值， 那么如果只有一个鸡蛋， 只能从最底层一层层的尝试
// 如果鸡蛋是无限个， 则使用二分的算法 ，这样是最快得出结果的， 但是如果 鸡蛋是大于 1 但是不是无限个数的时候
// 我们怎么样选择才是最好的选择呢.

func main(){
    fmt.Println(superEggDrop(1, 2))
}

//K 个鸡蛋 N 层楼
func superEggDrop(K int, N int) int {

    step := 0
    for  {
        if K <= 0 || N <= 1{
            break
        }
        center := 0
        if (N - 0) % 2 == 0 {
            center = (N - 1) / 2
        } else {
            center = N / 2 - 1
        }
        N = center - 1
        step++
        K--
    }
    return step
}