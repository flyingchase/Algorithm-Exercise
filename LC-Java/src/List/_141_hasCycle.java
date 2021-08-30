package List;

import DataStructure.ListNode;

public class _141_hasCycle {

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
        System.out.println(new _141_hasCycle().hasCycle(p1));
    }
    public boolean hasCycle(ListNode head) {

        if (head==null||head.next==null) {
            return false;
        }

        ListNode fast=head,slow=head;
        while (fast != null || fast.next != null) {
            fast=fast.next.next;
            slow=slow.next;
            if (slow ==fast) {
                return true;
            }
        }
        return  false;
    }
}
