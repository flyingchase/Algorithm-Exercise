package HotReview;

import DataStructure.ListNode;

/**
 * @author :qlzhou;
 * @date: :创建于2021/8/3010:49 上午
 */
public class E_206_reverseListNode {

    public ListNode reverseList(ListNode head) {
        ListNode cur = head;
        ListNode prev = null;
        prev.next=head;
        while (cur!=null) {
            ListNode next = cur.next;
            cur.next=prev;
            prev=cur;
            cur=cur.next;
        }
        return prev.next;
    }

}
