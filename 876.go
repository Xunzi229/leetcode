//https://leetcode-cn.com/problems/middle-of-the-linked-list/

package main

/*
type ListNode struct {
    Val int
    Next *ListNode
}
*/

// 利用快慢指针很好解决寻找中间数的问题
func middleNode(head *ListNode) *ListNode {
    quick := head

    for quick != nil && quick.Next != nil{
        quick, head = quick.Next.Next, head.Next
    }

    return head
}