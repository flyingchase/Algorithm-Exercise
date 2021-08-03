package List;

import DataStructure.ListNode;

import java.util.PriorityQueue;

public class _21_mergetwoLists {

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(3);
        ListNode p3 = new ListNode(5);
        ListNode p4 = new ListNode(7);
        ListNode p5 = new ListNode(2);
        ListNode p6 = new ListNode(4);
        ListNode p7 = new ListNode(6);
        ListNode p8 = new ListNode(8);
        ListNode p9 = new ListNode(10);

        p1.next=p2;
        p2.next=p3;
        p3.next=p4;

        // p4.next=p5;
        p5.next=p6;
        p6.next=p7;
        // 环
        p7.next=p8;

        ListNode listNode = new _21_mergetwoLists().mergeTwoLists1(p1, p5);
        while (listNode != null) {
            System.out.println(listNode.val);
            listNode=listNode.next;
        }
    }
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1==null) {
            return l2;
        }
        if (l2==null) {
            return l1;
        }

        ListNode p1 =l1,p2 = l2,cur=new ListNode(-1);
        ListNode root=cur;
        while (p1 != null && p2 != null) {
            if (p1.val<p2.val) {

                cur.next=p1;
                p1=p1.next;
            } else {
                cur.next=p2;
                p2=p2.next;
            }
            cur=cur.next;
        }
        if (p1 != null) {
            cur.next=p1;
        }
        if (p2!=null) {
            cur.next=p2;
        }

        return  root.next;

    }

    public ListNode mergeTwoLists1(ListNode l1, ListNode l2) {
        PriorityQueue<ListNode> queue = new PriorityQueue<>((ListNode list1, ListNode list2) -> (list2.val - list1.val));

        /*judge the legalness of  imports*/
        ListNode root =l1;
        while (l1.next != null) {
            l1=l1.next;
        }
        l1.next=l2;

        ListNode cur= root;
        while (cur != null) {
            queue.add(cur);
            cur=cur.next;
        }
        ListNode dummy02 = new ListNode(-1);
        dummy02.next=root;
        while (dummy02.next!=null) {
            deleteNode(dummy02);
        }


        int size = queue.size();
        ListNode dummy = new ListNode(-1);
        ListNode lastNode =dummy;
        while (!queue.isEmpty()) {
            lastNode.next = queue.poll();
            // assert lastNode != null;
            lastNode=lastNode.next;
            if (lastNode.next!=null) {
                queue.add(lastNode.next);
            }
        }
        return dummy;


    }

    private  void deleteNode (ListNode node ) {
        node.next=node.next.next;
    }

}
