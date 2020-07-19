//https://leetcode-cn.com/problems/reorder-list/
package reorder_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := findMiddle(head)
	left := head
	right := mid.Next
	mid.Next = nil

	right = reverse(right)
	for right != nil {
		lNext := left.Next
		rNext := right.Next
		left.Next = right
		right.Next = lNext
		left = lNext
		right = rNext
	}
}

// 快慢指针寻找中位数
func findMiddle(head *ListNode) *ListNode {
	slow, quick := head, head.Next

	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	return slow
}

// 翻转链表
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
