package List;

import DataStructure.ListNode;

public class HasRingForList {
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

        ListNode head1 = new ListNode(0);
        ListNode p01 = new ListNode(1);
        ListNode p02 = new ListNode(2);
        ListNode p03 = new ListNode(3);
        ListNode p04 = new ListNode(4);
        ListNode p05 = new ListNode(5);
        ListNode p06 = new ListNode(6);
        ListNode p07 = new ListNode(7);
        ListNode p08 = new ListNode(8);
        ListNode p09 = new ListNode(9);
        head.next = p1;
        p1.next = p2;
        p2.next = p3;
        p3.next = p4;
        p4.next = p5;
        p5.next = p6;
        p6.next = p7;
        p7.next = p8;
        p8.next = p9;

        head1.next = p01;
        p01.next = p02;
//        p02.Next = p03;
//        p03.Next = p04;
//        p04.Next = p05;
//        p02.Next = p06;
        p02.next = p7;

//        p06.Next = p7;
//        p7.Next=p4;

//        ListNode res = new HasRingForList().hasRingForList(head);
//        ListNode res1 = new HasRingForList().hasRingForList(head1);
//        if (res != null) {
//
//            System.out.println(res.val);
//        }else {
//            System.out.println("NULL");
//        }

        {
            int cnt = 0, cnt1 = 0;
            ListNode cur = head, cur1 = head1;
            while (cur != null) {
                cnt++;
                cur = cur.next;

            }
            while (cur1 != null) {
                cnt1++;
                cur1 = cur1.next;
            }

            int gap = Math.abs(cnt - cnt1);
            while (gap-- > 0) {
                if (cnt > cnt1) {
                    head = head.next;
                } else {
                    head1 = head1.next;
                }
            }
            while (head!=head1) {
                head = head.next;
                head1=head1.next;
            }
            System.out.println(head.val);
        }


    }

    public ListNode hasRingForList(ListNode head) {
        if (head == null) {
            return null;
        }
        ListNode fast = head, slow = head;
        do {
            fast = fast.next.next;
            slow = slow.next;
        } while (fast != slow && fast != null && slow != null);
        if (fast != slow) {
            return null;
        }
        fast = head;
        while (fast != slow) {
            fast = fast.next;
            slow = slow.next;
        }
        return slow;

    }


}
