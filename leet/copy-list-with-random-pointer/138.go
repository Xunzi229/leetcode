package copy_list_with_random_pointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 将会存储每个原节点对应的副本节点
	lists := make(map[*Node]*Node)

	// 每个原Random节点对应的 存储原Random节点地址的 副本节点
	lists2 := make(map[*Node]*Node)

	newHead := new(Node)
	newHead = &Node{Next: head}

	tmpHead := newHead
	for head != nil {
		tmpHead.Next = &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		}

		tmpHead = tmpHead.Next

		// 存这个值表示 后面 如果 有random指向这个节点的时候 ， 更改成指向这个节点的副本节点
		// 对应的是1
		lists[head] = tmpHead

		// 存这个值表示， 遍历到后面的时候 ，如果后面的节点存在， 则将这个节点的Random指向后面的副本节点
		// 此时是反向的存储。 表示到这个节点的时候需要将其指向这个节点的副本指向副本
		// 对应的是 2
		if head.Random != nil {
			lists2[head.Random] = tmpHead
		}

		// 接下来 我们考虑的是 他有没有指向别的节点， 或者别的节点有没有指向他

		// 当前节点指向前面的节点  ===> 1
		if lists[head.Random] != nil {
			(*tmpHead).Random = lists[head.Random]
		}

		// 前面的节点指向当前的节点 ==> 2
		if lists2[head] != nil {
			lists2[head].Random = tmpHead
		}

		head = head.Next
	}

	return newHead.Next
}
