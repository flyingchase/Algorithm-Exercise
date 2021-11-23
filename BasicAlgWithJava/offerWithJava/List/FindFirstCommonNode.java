package List;

import InnerStruct.ListNode;

public class FindFirstCommonNode {
    public static ListNode findFirstCommListNode(ListNode l1, ListNode l2) {
        ListNode p1 = l1, p2 = l2;
        while (l1 != l2) {
            l1 = (l1 == null) ? p2 : l1.next;
            l2 = (l2 == null) ? p1 : l2.next;

        }
        return l1;

    }

    public static void main(String[] args) {
        ListNode head1 = new ListNode(1);

    }
}
