// https://leetcode-cn.com/problems/partition-list/
package partition_list

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

//func partition(head *ListNode, x int) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	tmpHead := &ListNode{Val: 0}
//	bigLists := make([]*ListNode, 0)
//	lowLists := make([]*ListNode, 0)
//
//	for head != nil {
//		if head.Val >= x {
//			bigLists = append(bigLists, head)
//		} else {
//			lowLists = append(lowLists, head)
//		}
//		head = head.Next
//	}
//
//	tmp := tmpHead
//	for _, v := range lowLists {
//		v.Next = nil
//		tmpHead.Next = v
//		tmpHead = tmpHead.Next
//	}
//
//	for _, v := range bigLists {
//		tmpHead.Next = v
//		v.Next = nil
//		tmpHead = tmpHead.Next
//	}
//
//	return tmp.Next
//}

// 其实使用哑巴节点更符合逻辑吧
// 将所有的大于该数的节点都 存放到另一个链表中

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	// 创建两个空节点
	headAssume := &ListNode{Val: 0}
	tailAssume := &ListNode{Val: 0}

	tail := tailAssume

	headAssume.Next = head
	head = headAssume

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			t := head.Next
			head.Next = head.Next.Next

			// 将大于这个数的节点放到另一个链表中
			tail.Next = t
			tail = tail.Next
		}
	}

	// 将链表的最后指向置空
	tail.Next = nil

	// 拼接两个节点
	head.Next = tailAssume.Next

	return headAssume.Next
}
