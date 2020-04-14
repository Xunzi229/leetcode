// https://leetcode-cn.com/problems/add-two-numbers-ii/

package main

/**
 * Definition for singly-linked list.
 //* type ListNode struct {
 //*     Val int
 //*     Next *ListNode
 //* }
 */

type ListNode1 struct {
    Val int
    Next *ListNode1
}

func addTwoNumbers2(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
    list1Array := make([]int, 0)
    list2Array := make([]int, 0)
    list3Array := make([]int, 0)
    for l1 != nil {
        list1Array = append(list1Array, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        list2Array = append(list2Array, l2.Val)
        l2 = l2.Next
    }
    a1Len := len(list1Array)
    a2Len := len(list2Array)

    upPlus := 0
    if a1Len >= a2Len {
        tmpI := a2Len - 1
        for i := len(list1Array) - 1; i >= 0; i-- {
            if tmpI < 0 {
                if upPlus + list1Array[i] >= 10 {
                    list3Array = append(list3Array, (upPlus + list1Array[i]) % 10)
                    upPlus = 1
                } else {
                    list3Array = append(list3Array, upPlus + list1Array[i])
                    upPlus = 0
                }
            } else {
                if upPlus + list1Array[i] + list2Array[tmpI] >= 10 {
                    list3Array = append(list3Array, (upPlus + list1Array[i] + list2Array[tmpI]) % 10)
                    upPlus = 1
                } else {
                    list3Array = append(list3Array, upPlus + list1Array[i] + list2Array[tmpI])
                    upPlus = 0
                }
                tmpI--
            }

        }
    } else {
        tmpI := a1Len - 1
        for i := len(list2Array) - 1; i >= 0; i-- {
            if tmpI < 0 {
                if upPlus + list2Array[i] >= 10 {
                    list3Array = append(list3Array, (upPlus + list2Array[i]) % 10)
                    upPlus = 1
                } else {
                    list3Array = append(list3Array, upPlus + list2Array[i])
                    upPlus = 0
                }
            } else {
                if upPlus + list2Array[i] + list1Array[tmpI] >= 10 {
                    list3Array = append(list3Array, (upPlus + list2Array[i] + list1Array[tmpI]) % 10)
                    upPlus = 1
                } else {
                    list3Array = append(list3Array, upPlus + list2Array[i] + list1Array[tmpI])
                    upPlus = 0
                }
                tmpI--
            }

        }
    }
     if upPlus == 1 {
         list3Array = append(list3Array, 1)
     }
    head := new(ListNode1)
    currentNode := head
    for i := len(list3Array) - 1; i >= 0; i-- {
        if head.Val == 0 {
            if list3Array[i] == 0 {
                continue
            }
            head.Val = list3Array[i]
            continue
        }
        currentNode.Next =  &ListNode1{
            Val: list3Array[i],
        }
        currentNode = currentNode.Next
    }
    return head
}