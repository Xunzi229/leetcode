// https://leetcode-cn.com/problems/add-two-numbers/

package main
import (
    "fmt"
    "strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){
    a, b := []int{9,9}, []int{9}

    l1, l2, tmp  := new(ListNode), new(ListNode), new(ListNode)

    for i := 0; i < len(a); i++{
        if i == 0 {
            l1.Val = a[i]
            tmp = l1
        } else {
            tmp.Next = &ListNode{Val: a[i]}
            tmp = tmp.Next
        }
    }

    for i := 0; i < len(b); i++{
        if i == 0 {
            l2.Val = b[i]
            tmp = l2
        } else {
            tmp.Next = &ListNode{Val: b[i]}
            tmp = tmp.Next
        }
    }

    list3 := addTwoNumbers (l1, l2)
    nextNode := list3

    fmt.Println(a, b)
    fmt.Println("完成---开始打印")

    for true {
        fmt.Printf("%d -> ", nextNode.Val)
        if nextNode.Next == nil {
            break
        }
        nextNode = nextNode.Next
    }
    return
}

func appReverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func addTwoNumbers (l1 *ListNode, l2 *ListNode) *ListNode{
    start1, start2 := l1, l2

    for true {
       if l1 == nil {
           return start2
       }
       if l2 == nil {
           return start1
       }

        tmp := l1
        var totalVal int
        var t string
        totalVal = l1.Val + l2.Val

        t = strconv.Itoa(totalVal)
        if totalVal >= 10 {
            l2.Val, _ = strconv.Atoi(string(t[1]))
            l1.Val, _ = strconv.Atoi(string(t[1]))

            if l1.Next != nil {
                appRangAdd(l1.Next)
            } else {
                l1.Next = &ListNode{Val: 1}
            }
            l1 = tmp
        } else {
            l1.Val, _ = strconv.Atoi(string(t[0]))
            l2.Val, _ = strconv.Atoi(string(t[0]))
        }

        l1 = l1.Next
        l2 = l2.Next
    }
    return start1
}

func appRangAdd(num *ListNode){
    if num.Val + 1 >= 10 {
        t := strconv.Itoa(num.Val + 1)
        num.Val, _ = strconv.Atoi(string(t[0]))
        if num.Next == nil {
            num.Next = &ListNode{Val: 1}
            return
        } else {
            appRangAdd(num.Next)
        }
    } else {
        num.Val = num.Val + 1
    }
}