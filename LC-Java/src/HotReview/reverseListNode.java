package HotReview;

import DataStructure.ListNode;

/**
 * @author :qlzhou;
 * @date: :创建于2021/8/3010:49 上午
 */
public class reverseListNode {
    // reverse linkedlist all
    public ListNode reverseList(ListNode head) {
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

    // reverse linkedlist from index left to right ; assert 1<=left<=right<=len
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (head == null || left == right) {
            return head;
        }
        ListNode dummy = new ListNode(-1);
        dummy.next = head;

        ListNode cur = dummy;
        for (int i = 0; i < left - 1; i++) {
            cur = cur.next;
        }
        // ensure wall 为逆转区间的左侧前一个节点 包括 left=head 时 wall即为 dummy
        ListNode wall = cur;
        ListNode leftNode = wall.next, rightNode = leftNode.next;
        // 头插法逐个逆转
        for (int j = left+1; j <= right; j++) {
            leftNode.next=rightNode.next;
            rightNode.next=wall.next;
            wall.next=rightNode;
            rightNode=leftNode.next;
        }
        // return the ture head
        return dummy.next;

    }
}
