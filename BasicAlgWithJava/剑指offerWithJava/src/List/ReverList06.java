package List;

import DataStructure.ListNode;

import java.util.ArrayList;
import java.util.Iterator;

@SuppressWarnings({"all"})
public class ReverList06 {
    public static  ArrayList<Integer> printListFromTailToHead01(ListNode listNode) {
        ArrayList<Integer> res = new ArrayList();
        if (listNode==null) {
            return res;
        }
        recurively(res,listNode);
        return res;
    }
    public static void recurively(ArrayList<Integer> res, ListNode node) {
        if (node==null) {
            return;
        }
        recurively(res,node.Next);
        res.add(node.val);
        return;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(-1);
        head.Next=new ListNode(1);
        head.Next.Next=new ListNode(2);
        head.Next.Next.Next=new ListNode(3);
        head.Next.Next.Next.Next=new ListNode(4);
        head.Next.Next.Next.Next.Next=new ListNode(5);
        head.Next.Next.Next.Next.Next.Next=new ListNode(6);

        ArrayList<Integer> res = new ArrayList<>();
//        res= printListFromTailToHead01(head);
        res= (ArrayList<Integer>) printListFromTailToHead02(head);
        for (Object o:
             res) {
            System.out.println(o);
        }

        Iterator ite = res.iterator();
        while (ite.hasNext()) {
            Object next = ite.next();
            System.out.println(next);
        }

    }
    public static ArrayList<? extends Object> printListFromTailToHead02(ListNode listNode) {
        ArrayList<Object> res = new ArrayList<>();
        if (listNode==null) {
            return res;
        }
        return unrecursively(listNode);
    }

    public static ArrayList<Integer> unrecursively(ListNode node) {
        ArrayList<Integer> res = new ArrayList<>();
        ListNode newHead = reversedNode(node);
        ListNode p =newHead;
        while (p != null) {
            res.add(p.val);
            p=p.Next;
        }
        reversedNode(newHead);
        return res;
    }

    public static ListNode reversedNode(ListNode node) {
        ListNode pre =null,cur=node,next;
        if (cur!=null) {
            next = cur.Next;
            cur.Next=pre;
            pre=cur;
            cur=next;
        }
        return pre;
    }
}
