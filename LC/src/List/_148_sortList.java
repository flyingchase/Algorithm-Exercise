package List;

import DataStructure.ListNode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.LinkedList;

public class _148_sortList {

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(9);
        ListNode p2 = new ListNode(6);
        ListNode p3 = new ListNode(4);
        ListNode p4 = new ListNode(3);
        ListNode p5 = new ListNode(8);
        ListNode p6 = new ListNode(7);
        ListNode p7 = new ListNode(1);
        ListNode p8 = new ListNode(5);
        ListNode p9 = new ListNode(2);
head.next=p1;
        p1.next=p2;
        p2.next=p3;
        p3.next=p4;
        p4.next=p5;
        p5.next=p6;
        p6.next=p7;
        // 环
        p7.next=p8;
        p8.next = p9;
        ListNode newHead = new _148_sortList().sortList(head);
        while (newHead != null) {
            System.out.println(newHead.val);
            newHead=newHead.next;
        }
    }
    public ListNode sortList(ListNode head) {

        ListNode head1 = head;
        ArrayList<Integer> lists = new ArrayList<Integer>();
        int cnt = 0,i=0;
        while (head1 != null) {
            lists.add(head1.val);
            head1 = head1.next;
            cnt++;
        }
        int[] nums = new int[cnt];

        // Comparator.comparing(lists,(o1, o2)->(o2-o1));
        for (Integer list : lists) {
            // System.out.println(list);
            nums[i++]=list;

        }
        Arrays.sort(nums);
        ListNode newHead = new  ListNode(-1);
        ListNode cur =newHead;
        for (int num : nums) {
            ListNode p = new ListNode(num);
            cur.next=p;
            cur=cur.next;
        }

        return newHead.next;
    }

    /*private void quicksot(ListNode head, ListNode end) {
        if (head.val==end.val) {
            return;
        }
        swap(head, end);

        partition(head, end);


    }

    private ListNode partition(ListNode head, ListNode end) {
        ListNode less = new ListNode(-1);
        ListNode more= end;
        less.next=head;
        while (head.val<more.val) {
            if (head.val < end.val) {
                swap(less.next,end);
            }else if (head.val>end.val){
                swap(head,);
            }
        }
    }

    private void swap(ListNode head, ListNode end) {
        int temp = head.val;
        head.val = end.val;
        end.val = temp;
    }
*/

}
