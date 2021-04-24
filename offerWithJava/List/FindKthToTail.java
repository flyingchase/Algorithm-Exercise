 package List;

import Function.*;
import InnerStruct.*;

public class FindKthToTail {
    public static ListNode findKthToTail(ListNode head, int k) {
        if(head==null) {
            return null;
        }
        ListNode p1=head;
        while (p1!=null&&k-->0) {
            p1=p1.next;
        }
        if(k>0) {
            return null;
        }
        ListNode p2=head;
        while (p1!=null) {
            p1=p1.next;
            p2=p2.next;
        }
        return p2;
    }

    public static void main(String[] args) {
        int[] nums={9,8,7,6,5,4,3,2,1};
        ListOperate listOperate = new ListOperate();
        ListNode head = listOperate.arrayToListNode(nums);
        listOperate.printListNode(head);
        head=findKthToTail(head, 1);
        listOperate.printListNode(head);

    }
}
