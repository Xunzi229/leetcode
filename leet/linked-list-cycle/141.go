package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, quick := head, head

	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			return true
		}
	}
	return false
}
