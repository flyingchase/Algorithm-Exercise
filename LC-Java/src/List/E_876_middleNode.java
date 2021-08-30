package List;

import DataStructure.ListNode;

/*给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。*/
public class E_876_middleNode {
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

        // p1.next=p2;
        // p2.next=p3;
        p3.next=p4;
        p4.next=p5;
        p5.next=p6;
        p6.next=p7;
        p7.next=p8;
        System.out.println(new E_876_middleNode().middleNode(p1).val);
        // 环
        // p7.next=p4;
    }
    public ListNode middleNode(ListNode head) {
        ListNode cur = head;
        int count =0;
        while (cur != null) {
            count++;
            cur=cur.next;
        }
        int mid = count/2;
        while (mid-- > 0) {
            head=head.next;
        }
        return  head;
    }
}
