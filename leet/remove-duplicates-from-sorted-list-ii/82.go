// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
package remove_duplicates_from_sorted_list_ii

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmpHead := head
	for head.Next != nil {
		for head.Next.Val == head.Val {
			head.Next = head.Next.Next
			if head.Next == nil {
				return tmpHead
			}
		}
		head = head.Next
	}

	return tmpHead
}
