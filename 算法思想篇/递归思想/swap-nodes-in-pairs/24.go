/*https://leetcode-cn.com/problems/swap-nodes-in-pairs/*/
package swap_nodes_in_pairs


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    tmp := &ListNode{}
    tmp.Next = head
    reverse(tmp)
    return tmp.Next
}

// 在交换的过程中， 保证下一次的循环能连接上
func reverse(node *ListNode) {
    if node == nil || node.Next == nil {
        return
    }

    if node.Next != nil && node.Next.Next != nil {
        tmpHead := node.Next.Next
        node.Next.Next.Next, node.Next.Next = node.Next, node.Next.Next.Next
        node.Next = tmpHead
    }
    reverse(node.Next.Next)
}