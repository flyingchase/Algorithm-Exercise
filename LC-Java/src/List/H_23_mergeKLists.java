package List;

import DataStructure.ListNode;

import java.util.Collections;
import java.util.Comparator;
import java.util.PriorityQueue;

/*给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。*/
public class H_23_mergeKLists {
    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(2);
        ListNode p3 = new ListNode(3);
        ListNode p4 = new ListNode(4);
        ListNode p5 = new ListNode(5);
        ListNode p6 = new ListNode(6);
        ListNode p7 = new ListNode(7);
        ListNode p8 = new ListNode(8);
        ListNode p9 = new ListNode(9);

        p1.next = p2;
        p2.next = p3;
        // p3.next=p4;
        p4.next = p5;
        p5.next = p6;
        // p6.next=p7;
        // 环
        p7.next = p8;
        p8.next = p9;

        // ListNode[] lists= new ListNode[]{p1,p4,p7};
        // ListNode node = new H_23_mergeKLists().mergeKLists(lists);
        // while (node != null) {
        //     System.out.println(node.val);
        //     node=node.next;
        // }
        new H_23_mergeKLists().mergeKLists(new ListNode[]{});

    }


    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null) {
            return null;
        }
        PriorityQueue<ListNode> pq = new PriorityQueue<>(new Comparator<ListNode>() {
            @Override
            public int compare(ListNode o1, ListNode o2) {
                return o1.val - o2.val;
            }
        });


        Collections.addAll(pq, lists);
        // for (ListNode list : lists) {
        //     if (list!=null) {
        //         pq.add(list);
        //     }
        // }

        ListNode dummy = new ListNode(-1);
        ListNode cur = dummy;
        // dummy.next=cur;
        while (!pq.isEmpty()) {
            ListNode node = pq.poll();
            cur.next = node;
            cur = cur.next;
            // node为每个单链表的头结点 要把后续结点加入其中
            if (cur.next != null) {
                pq.add(cur.next);
            }
        }

        return dummy.next;

    }


    // ues merge
    /*
    public ListNode mergeKLists(ListNode[] lists) {
	if (lists == null || lists.length == 0)
		return null;
    return mergeKLists(lists, 0, lists.length - 1);
}
private ListNode mergeKLists(ListNode[] lists, int start, int end) {
	if (start == end) {
		return lists[start];
	} else if (start < end){
		int mid = (end - start) / 2 + start;
		ListNode left = mergeKLists(lists, start, mid);
		ListNode right = mergeKLists(lists, mid + 1, end);
		return mergeTwoLists(left, right);
	} else {
		return null;
	}
}
*/
}
