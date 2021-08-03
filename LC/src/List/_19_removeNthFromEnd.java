package List;

import DataStructure.ListNode;

/*给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

 */
public class _19_removeNthFromEnd {

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(2);
        ListNode p3 = new ListNode(3);
        ListNode p4 = new ListNode(4);
        ListNode p5 = new ListNode(5);
        ListNode p6 = new ListNode(6);
        ListNode p7 = new ListNode(7);
        ListNode p8 = new ListNode(8);
        ListNode p9 = new ListNode(9);

        p1.next = p2;
        p2.next = p3;
        p3.next = p4;
        p4.next = p5;
        p5.next = p6;
        p6.next = p7;
        // 环
        // p7.Next=p4;

        ListNode root = new _19_removeNthFromEnd().removeNthFromEnd(p8, 1);
        while (root != null) {
            System.out.println(root.val);
            root = root.next;
        }

    }

    public ListNode removeNthFromEnd(ListNode head, int n) {

        ListNode newHead = new ListNode(-1);
        newHead.next = head;
        ListNode p1 = newHead, p2 = newHead;

        for (int i = 0; i < n; i++) {
            p1 = p1.next;
        }
        while (p1.next != null) {
            p1 = p1.next;
            p2 = p2.next;
        }
        p2.next = p2.next.next;

        return newHead.next;
    }
}
