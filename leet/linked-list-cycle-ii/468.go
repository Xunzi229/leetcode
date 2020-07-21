package linked_list_cycle_ii

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

func detectCycle(head *ListNode) *ListNode {
	slow, quick := head, head

	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			quick = head

			// 注意这个地方， 如果第一次相遇的点就是起始点， 则说明， 这个点就是环的起始点
			if quick == slow {
				return quick
			}

			for {
				quick = quick.Next
				slow = slow.Next
				if quick == slow {
					return slow
				}
			}
		}

	}
	return nil
}
