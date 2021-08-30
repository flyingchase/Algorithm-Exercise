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
        while (cur!=null) {
            ListNode next = cur.next;
            cur.next=prev;
            prev=cur;
            cur=next;
        }
        return prev;
    }

    // reverse linkedlist from index left to right ; assert 1<=left<=right<=len
    public ListNode reverseBetween(ListNode head, int left, int right) {
        
    }
}
