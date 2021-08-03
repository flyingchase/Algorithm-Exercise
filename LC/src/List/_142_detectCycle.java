package List;

import DataStructure.ListNode;

public class _142_detectCycle {
    boolean falg =false;

    public static void main(String[] args) {

    }

    public ListNode detectCycle(ListNode head) {
        if (!falg) {
            return null;
        }
        if (falg) {
            ListNode p1 = hasCycle(head);
            ListNode p2=head;
            int cnt=0;
            while (p1 != p2) {
                p1=p1.next;
                p2=p2.next;
                cnt++;
            }
            return p1;

        }
        return null;
    }

    private ListNode hasCycle(ListNode head) {
        ListNode fast = head, slow = head;
        while (fast != null||fast.next != null) {
            fast=fast.next;
            slow=slow.next;
            if (fast==slow) {
                falg=true;
                return slow;
            }
        }
        return null;
    }
}
