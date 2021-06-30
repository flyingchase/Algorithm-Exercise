package List;

import Function.ListOperate;
import InnerStruct.ListNode;

/* 
输入一个链表，反转链表后，输出新链表的表头。

###  */
public class ReverseList {
    // 递归
    public static ListNode reversListNode(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode next = head.next;
        head.next = null;
        ListNode newHead = reversListNode(next);
        next.next = head;
        return newHead;

    }

    // 头插法
    public static ListNode reversListNode02(ListNode head) {
        ListNode newList = new ListNode(-1);
        while (head != null) {
            ListNode next = head.next;
            head.next = newList.next;
            newList.next = head;
            head = next;
        }
        return newList.next;
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        ListOperate listOperate = new ListOperate();
        ListNode head = listOperate.arrayToListNode(nums);
        head = reversListNode(head);
        listOperate.printListNode(head);

        head = reversListNode02(head);
        listOperate.printListNode(head);

    }
}
