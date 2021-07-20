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
        head.Next = p1;
        p1.Next = p2;
        p2.Next = p3;
        p3.Next = p4;
        p4.Next = p5;
        p5.Next = p6;
        p6.Next = p7;
        p7.Next = p8;
        p8.Next = p9;

        head1.Next = p01;
        p01.Next = p02;
//        p02.Next = p03;
//        p03.Next = p04;
//        p04.Next = p05;
//        p02.Next = p06;
        p02.Next = p7;

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
                cur = cur.Next;

            }
            while (cur1 != null) {
                cnt1++;
                cur1 = cur1.Next;
            }

            int gap = Math.abs(cnt - cnt1);
            while (gap-- > 0) {
                if (cnt > cnt1) {
                    head = head.Next;
                } else {
                    head1 = head1.Next;
                }
            }
            while (head!=head1) {
                head = head.Next;
                head1=head1.Next;
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
            fast = fast.Next.Next;
            slow = slow.Next;
        } while (fast != slow && fast != null && slow != null);
        if (fast != slow) {
            return null;
        }
        fast = head;
        while (fast != slow) {
            fast = fast.Next;
            slow = slow.Next;
        }
        return slow;

    }


}
