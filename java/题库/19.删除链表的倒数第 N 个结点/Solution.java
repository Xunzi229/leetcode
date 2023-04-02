/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
*/

public ListNode removeNthFromEnd(ListNode head, int n) {
    List<ListNode> list = new ArrayList<>();

    while (head.next != null) {
        list.add(head);
        head = head.next;
    }
    int nx = list.size() - n;

    System.out.println(nx);
    ListNode newHead = new ListNode();

    ListNode currentHead = newHead;
    for (int index = 0; index < list.size(); index++) {
        if (index == nx) {
            continue;
        }
        currentHead.next = list.get(index);
        list.get(index).next = null;
        currentHead = currentHead.next;
    }
    return newHead.next;
}

public ListNode removeNthFromEnd1(ListNode head, int n) {
    ListNode dummy = new ListNode(0);
    dummy.next = head;
    ListNode fast = dummy;
    ListNode slow = dummy;
    for (int i = 0; i <= n; i++) {
        fast = fast.next;
    }
    while (fast != null) {
        fast = fast.next;
        slow = slow.next;
    }
    slow.next = slow.next.next;
    return dummy.next;
}