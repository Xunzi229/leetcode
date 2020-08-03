//https://leetcode-cn.com/problems/merge-k-sorted-lists/

package merge_k_sorted_lists

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

func mergeKLists(lists []*ListNode) *ListNode {
	result := &ListNode{}

	for _, l := range lists {
		r := result.Next

		tmp := result
		for l != nil && r != nil {
			if l.Val > r.Val {
				tmp.Next = r
				r = r.Next
			} else {
				tmp.Next = l
				l = l.Next
			}
			tmp = tmp.Next
		}
		if l != nil {
			tmp.Next = l
		}

		if r != nil {
			tmp.Next = r
		}
	}
	return result.Next
}

/*
func mergeKLists(lists []*ListNode) *ListNode {
	headList := new(ListNode)
	currentList := headList

	for len(lists) != 0 {
		if lists[0] == nil {
			lists = lists[1:]
			continue
		}

		minValue := lists[0].Val
		siteValue := 0

		tmpNilSite := make([]int, 0)
		for i, v := range lists {
			if v == nil {
				tmpNilSite = append(tmpNilSite, i)
				continue
			}
			if v.Val < minValue {
				siteValue = i
				minValue = v.Val
			}
		}
		currentList.Next = &ListNode{
			Val: minValue,
		}
		currentList = currentList.Next

		if lists[siteValue].Next == nil {
			tmpNilSite = append(tmpNilSite, siteValue)
		} else {
			lists[siteValue] = lists[siteValue].Next
		}
		sort.Ints(tmpNilSite)

		preI := 0
		for _, v := range tmpNilSite {
			_v := v - preI
			lists = append(lists[:_v], lists[_v+1:]...)
			preI += 1
		}
	}

	return headList.Next
}
*/
