package List;

import DataStructure.ListNode;

public class _206_reverseList {
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

        p1.next=p2;
        p2.next=p3;
        p3.next=p4;
        p4.next=p5;
        p5.next=p6;
        p6.next=p7;
        // 环
        // p7.next=p4;
        ListNode reverseList = new _206_reverseList().reverseList(p1);
        while (reverseList != null) {
            System.out.println(reverseList.val);
            reverseList=reverseList.next;
        }

    }
    public ListNode reverseList(ListNode head) {
        if (head==null) {
            return head;
        }
        ListNode prev = null,cur = head;

        while (cur != null) {
            ListNode next = cur.next;
            cur.next=prev;
            prev=cur;
            cur=next;
        }
        return prev;
    }

}
