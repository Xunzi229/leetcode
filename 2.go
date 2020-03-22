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
    a, b := []int{2,4,3, 12}, []int{5,6,4,33}

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

    l3 := addTwoNumbers(l1, l2)
    nextNode := l3

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

func appInsert(nums []int, site int) []int {
    numsLength := len(nums)
    rangTimes  :=0
    
    for true {
        if rangTimes > numsLength {
            break
        }
        if site > numsLength - 1 {
            nums = append([]int{1}, nums...)
            return nums
        }
        if nums[site] + 1 >= 10 {
            str := strconv.Itoa(nums[site] + 1)
            nums[site], _ = strconv.Atoi(string(str[1]))
            
            nums = appInsert(nums, site + 1)
        } else {
            nums[site] = site + 1
            return nums
        }
        rangTimes += 1
    }
    
    return nums
}

func appNum(aNum []int, bNum []int) (cNum []int) {
    rangeTimes, bMax, bCal := len(aNum), len(bNum), 0
    
    for i := 0; i <= rangeTimes - 1; i++ {
        x, y := aNum[i], 0
        if bCal < bMax {
            y = bNum[i]
        }
        if (x + y) >= 10 {
            str := strconv.Itoa(x + y)
            aNum[i], _ = strconv.Atoi(string(str[1]))
            
            aNum = appInsert(aNum, i + 1)
        } else {
            aNum[i] = x + y
        }
        cNum = append(cNum, aNum[i])
        bCal += 1
    }
    return cNum
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var nums1, nums2, nums3 []int

    tmpL1, tmpL2 := l1, l2
    for true  {
        nums1 = append(nums1, tmpL1.Val)
        if tmpL1.Next == nil {
            break
        }
        tmpL1 = tmpL1.Next
    }

    for true {
        nums2 = append(nums2, tmpL2.Val)
        if tmpL2.Next == nil {
            break
        }
        tmpL2 = tmpL2.Next
    }
    appReverse(nums1)
    appReverse(nums2)
    
    if len(nums1) >= len(nums2){
        nums3 = appNum(nums1, nums2)
    } else {
        nums3 = appNum(nums2, nums1)
    }
    appReverse(nums3)
    
    l3String := ""
    for i, _ := range nums3 {
       l3String += strconv.Itoa(nums3[i])
    }
    
    startL3, upNode := new(ListNode), new(ListNode)
    for i := len(l3String) - 1; i >= 0; i-- {
        t, _ := strconv.Atoi(string(l3String[i]))
        if i == len(l3String) - 1 {
           startL3.Val = t
           upNode = startL3
        } else {
           upNode.Next = &ListNode{Val: t}
           upNode = upNode.Next
        }
    }

    return startL3
}