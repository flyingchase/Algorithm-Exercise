/*
 * @Author: gunjianpan
 * @Date:   2021-05-08 17:21:59
 * @Last Modified by:   gunjianpan
 * @Last Modified time: 2021-05-08 19:29:27
 */
public class ReverseListNode {
    public static ListNode ReverseListNode(ListNode head) {
        ListNode cur=null,prev=null,next=null;
        cur=head;
        while (cur!=null) {
            next=cur.next;
            cur.next=prev;
            prev=cur;
            cur=next;

        }
        head=prev;
        return head;        
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(-1);
        ListNode l1 = new ListNode(0);
        head.next=l1;
        ListNode l2 = new ListNode(1);
        l1.next=l2;
        
        ListNode l3 = new ListNode(2);
        l2.next=l3;
        ListNode l4 = new ListNode(3);
        l3.next=l4;
        ListNode l5 = new ListNode(4);
        l4.next=l5;
        ListNode current=head;
        while (current!=null) {
            System.out.print(current.val+" ");
            current=current.next;
        }
        System.out.println();

        ReverseListNode(head);
        current=head;
        while (current!=null) {
            System.out.println(current.val+" ");
            current=current.next;
        }
        
    }
}
