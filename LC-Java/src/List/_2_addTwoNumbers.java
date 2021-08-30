package List;

import DataStructure.ListNode;

public class _2_addTwoNumbers {

    public static void main(String[] args) {
        ListNode head = new ListNode(4);
        ListNode p1 = new ListNode(4);
        ListNode p2 = new ListNode(3);
        ListNode p3 = new ListNode(1);
        ListNode p4 = new ListNode(1);
        ListNode p5 = new ListNode(5);
        ListNode p6 = new ListNode(2);
        ListNode p7 = new ListNode(7);
        ListNode p8 = new ListNode(8);
        ListNode p9 = new ListNode(9);


        p1.next = p2;
        p2.next = p3;
        // p3.Next = p4;

        p4.next = p5;
        p5.next = p6;

        ListNode listNode = new _2_addTwoNumbers().addTwoNumbers(p1, p4);
        while (listNode != null) {
            System.out.println(listNode.val);
            listNode = listNode.next;
        }
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {

        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l2;
        }
        ListNode cur1 = l1, cur2 = l2, root = l1;
        ListNode head1 = reverseListNode(l1);
        ListNode head2 = reverseListNode(l2);
        ListNode head = new ListNode(0);
        ListNode p = head;
        int value = 0;
        while (head1 != null && head2 != null) {
            value += head1.val + head2.val;
            head.next = new ListNode(value % 10);
            value /= 10;
            head = head.next;
            head1 = head1.next;
            head2 = head2.next;

        }
        while (head1 != null) {
            value += head1.val;
            head.next = new ListNode(value % 10);
            value /= 10;
            head = head.next;
            head1 = head1.next;
        }
        while (head2 != null) {
            value += head2.val;
            head.next = new ListNode(value % 10);
            value /= 10;
            head = head.next;
            head2 = head2.next;
        }
        if (value>0) {
            head.next = new ListNode(value);
        }
        return (p.next);
    }

    private ListNode reverseListNode(ListNode head) {
        ListNode cur = head;
        ListNode prev = null;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }
}
