//https://leetcode-cn.com/problems/sort-list/
package sort_list

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

func sortList(head *ListNode) *ListNode {
    return mergeSort(head)
}

// 使用归并排序
func mergeSort(head *ListNode) *ListNode {
  // 如果只有一个节点 直接就返回这个节点
  if head == nil || head.Next == nil{
      return head
  }

  // 找到中间的参数
  middle := findMiddle(head)
  tail := middle.Next
  middle.Next = nil

  // 开始归并
  left := mergeSort(head)
  right := mergeSort(tail)

  // 合并左右
  result := mergeResult(left, right)
  return result
}

func findMiddle(head *ListNode) *ListNode {
  slow, quick := head, head.Next

  for quick != nil && quick.Next != nil {
    quick = quick.Next.Next
    slow = slow.Next
  }
  return slow
}

func mergeResult(left, right *ListNode) *ListNode {
  current := &ListNode{Val: 0}

  head := current
  for left != nil && right != nil {
    if left.Val < right.Val {
      current.Next = left
      left = left.Next
    } else {
      current.Next = right
      right = right.Next
    }
    current = current.Next
  }

  if left != nil{
    current.Next = left
  }

  if right != nil{
    current.Next = right
  }

  return head.Next
}

