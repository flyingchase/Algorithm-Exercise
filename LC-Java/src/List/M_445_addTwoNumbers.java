package List;

import DataStructure.ListNode;

import java.util.Stack;

/*给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

*/

/*
*     public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        Stack<Integer> s1 = new Stack<Integer>();
        Stack<Integer> s2 = new Stack<Integer>();

        while(l1 != null) {
            s1.push(l1.val);
            l1 = l1.next;
        };
        while(l2 != null) {
            s2.push(l2.val);
            l2 = l2.next;
        }

        int sum = 0;
        ListNode list = new ListNode(0);
        while (!s1.empty() || !s2.empty()) {
            if (!s1.empty()) sum += s1.pop();
            if (!s2.empty()) sum += s2.pop();
            list.val = sum % 10;
            ListNode head = new ListNode(sum / 10);
            head.next = list;
            list = head;
            sum /= 10;
        }

        return list.val == 0 ? list.next : list;
    }
    * */
public class M_445_addTwoNumbers {
    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(2);
        ListNode p3 = new ListNode(3);
        ListNode p4 = new ListNode(4);
        ListNode p5 = new ListNode(9);
        ListNode p6 = new ListNode(9);
        ListNode p7 = new ListNode(7);
        ListNode p8 = new ListNode(8);
        ListNode p9 = new ListNode(9);

        // p1.next=p2;
        // p2.next=p3;
        // p3.next=p4;
        // // p4.next=p5;
        p5.next = p6;
        // p6.next=p7;
        // 环
        // p7.next=p4;
        ListNode listNode = new M_445_addTwoNumbers().addTwoNumbers(p1, p5);
        while (listNode != null) {
            System.out.println(listNode.val);
            listNode = listNode.next;
        }
        // ListNode reverse = new M_445_addTwoNumbers().reverseList(p1);
        // while (reverse != null) {
        //     System.out.println(reverse.val);
        //     reverse=reverse.next;
        // }
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }
        Stack<ListNode> s1 = new Stack<>();
        Stack<ListNode> s2 = new Stack<>();
        ListNode c1 = l1, c2 = l2;
        while (c1 != null) {
            s1.push(c1);
            c1 = c1.next;
        }
        while (c2 != null) {
            s2.push(c2);
            c2 = c2.next;
        }
        int sum = 0;

        ListNode head = new ListNode(-1);
        ListNode cur = head;
        while (!s1.isEmpty() && !s2.isEmpty()) {
            ListNode p1 = s1.pop();
            ListNode p2 = s2.pop();
            sum = sum + p1.val + p2.val;
            cur.next = new ListNode(sum % 10);
            sum /= 10;
            cur = cur.next;
        }
        while (!s1.isEmpty()) {
            if (sum != 0) {
                ListNode pop = s1.pop();
                pop.val += sum;
                sum /= 10;
                cur.next = new ListNode(pop.val);
                cur = cur.next;
                cur.next = null;
                break;
            }
            cur.next = s1.pop();
            cur.next.next = null;
            break;
        }
        while (!s2.isEmpty()) {
            while (sum != 0) {
                ListNode pop = s2.pop();
                sum += pop.val;
                cur.next = new ListNode(sum % 10);
                sum /= 10;
                cur = cur.next;
            }
            cur.next = s2.pop();
            cur.next.next = null;
            break;
        }
        while (sum != 0) {
            cur.next = new ListNode(sum % 10);
            cur = cur.next;
            sum /= 10;
        }
        cur = null;
        // return head.next;
        ListNode reverse = reverseList(head.next);
        return reverse;

    }


    private ListNode reverseList(ListNode head) {
        if (head == null) {
            return head;
        }
        ListNode prev = null, cur = head;

        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }

}
