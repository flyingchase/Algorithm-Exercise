/*
 * @Author: gunjianpan
 * @Date:   2021-05-08 19:50:19
 * @Last Modified by:   gunjianpan
 * @Last Modified time: 2021-05-08 19:58:57
 */

package ListNode;

public class ReverseListNode {
    public static ListNode reversListNode(ListNode head) {
        if (head==null||head.next==null) {
            return head;
        }

        ListNode cur=null,prev=null,next=null;
        cur=head;
        while (cur.next!=null) {
            next=cur.next;
            cur.next=prev;
            prev=cur;
            cur=next;
        }
        return cur;
    }

    public static void main(String[] args) {
        ListNode head =new ListNode(-1);
        ListNode l1= new ListNode(0);
        head.next=l1;
        ListNode l2= new ListNode(1);
        l1.next=l2;
        ListNode l3= new ListNode(2);
        l2.next=l3;
        ListNode l4= new ListNode(3);
        l3.next=l4;
        ListNode l5= new ListNode(4);
        l4.next=l5;
        ListNode current = head;
        while (current!=null) {
            System.out.print(current.val+" ");
            current=current.next;
        }
        System.out.println();
        reversListNode(head);
        while(head.next!=null) {
            System.out.print(head.val+" ");
            head=head.next;
        }
        System.out.println();
    }
}
