//https://leetcode-cn.com/problems/string-to-integer-atoi/

package main

import (
    "fmt"
)
/*
/*
 +-0123456789
32
43
45
48
49
50
51
52
53
54
55
56
57
*/


func main() {
    fmt.Println("0-1",                   myAtoi("0-1"))
    fmt.Println("-2147483649",           myAtoi("-2147483649"))
    fmt.Println("2147483649",            myAtoi("2147483649"))
    fmt.Println(" -0123456789",          myAtoi(" -0123456789"))
    fmt.Println("  0000000000012345678", myAtoi("  0000000000012345678"))
    fmt.Println("4193 with words",       myAtoi("4193 with words"))
    fmt.Println(".1",                    myAtoi(".1"))
    fmt.Println("42",                    myAtoi("42"))

    //fmt.Println(int(math.Pow10(i)), pow(10, i), i)
    //fmt.Println(pow(10, 3))
}

var (
  m = map[string]int{
      "0": 0,
      "1": 1,
      "2": 2,
      "3": 3,
      "4": 4,
      "5": 5,
      "6": 6,
      "7": 7,
      "8": 8,
      "9": 9,
  }
  max = pow(2, 31) - 1
  min = pow(2, 31) * -1
)

func myAtoi(str string) int {
    var newStr string
    var isStart, lt int = 0, 1

    for _, v := range str{
        if isStart == 0 && v == 32 {continue}
        if isStart == 0 && v == 43 {lt = 1; isStart += 1; continue}
        if isStart == 0 && v == 45 {lt = -1; isStart += 1; continue}
        if len(newStr) == 0 && v == 48 {isStart += 1; continue}

        if v < 48 || v > 57{break}
        newStr += string(v)
        isStart += 1
    }

    strLen := len(newStr)
    if strLen == 0 {return 0}
    if strLen > 10 {
        if lt == -1 {
          return min
        } else {
          return max
        }
    }

    var result = 0
    for i := len(newStr) - 1; i >= 0; i-- {
        _v := pow(10, i) * m[string(newStr[strLen - i - 1])]

        if lt == 1 && max - result <= _v {
            return max
        }
        if lt == -1 && max + 1 - result <= _v {
            return min
        }
        result += _v
    }
    return result * lt
}
func pow(x, y int) int {
    if y == 0 {return 1}
    result := x
    for i := 0; i < y - 1; i++{
        result *= x
    }
    return result
}