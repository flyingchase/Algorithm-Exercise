package List;

import DataStructure.ListNode;

import java.util.Comparator;
import java.util.HashMap;
import java.util.PriorityQueue;

/*存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

 */
public class E_83_deleteDuplicates {

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(1);
        ListNode p3 = new ListNode(2);
        ListNode p4 = new ListNode(3);
        ListNode p5 = new ListNode(3);
        ListNode p6 = new ListNode(5);
        ListNode p7 = new ListNode(7);
        ListNode p8 = new ListNode(7);
        ListNode p9 = new ListNode(9);

        p1.next = p2;
        p2.next = p3;
        p3.next = p4;
        p4.next = p5;
        // p5.next=p6;
        p6.next = p7;
        // 环
        p7.next = p8;

        ListNode listNode = new E_83_deleteDuplicates().deleteDuplicates(p1);
        while (listNode != null) {
            System.out.println(listNode.val);
            listNode = listNode.next;
        }
    }

    public ListNode deleteDuplicates(ListNode head) {
        PriorityQueue<ListNode> pq = new PriorityQueue<>(new Comparator<ListNode>() {
            @Override
            public int compare(ListNode o1, ListNode o2) {
                return o1.val - o2.val;
            }
        });

        ListNode curr = head, dummy = new ListNode(Integer.MAX_VALUE);
        dummy.next = head;
        ListNode cur = dummy;
        while (curr != null) {
            pq.add(curr);
            curr = curr.next;
        }

        while (!pq.isEmpty()) {
            ListNode poll = pq.poll();
            if (poll.val == cur.val) {
                continue;
            }
            cur.next = poll;
            cur = cur.next;
        }
        cur.next = null;
        return dummy.next;
    }

}
