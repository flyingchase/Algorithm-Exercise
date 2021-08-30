package List;

import DataStructure.ListNode;

public class _160_getIntersectionNode {


    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {

        if (headA == null || headB == null) {
            return null;
        }
        int gap = 0, lenA = 0, lenB = 0;
        ListNode pA = headA, pB = headB;
        while (pA != null) {
            lenA++;
            pA = pA.next;
        }
        while (pB != null) {
            pB = pB.next;
            lenB++;
        }
        gap = Math.abs(lenA - lenB);
        pA = headA;
        pB = headB;
        while (lenA > lenB && gap > 0) {
            gap--;
            pA = pA.next;
        }
        while (lenB > lenA && gap > 0) {
            gap--;
            pB=pB.next;
        }

        while (pA != null && pB != null) {
            if (pA==pB) {
                return pA;
            }
            pA = pA.next;
            pB=pB.next;
        }

        return null;


    }
}
