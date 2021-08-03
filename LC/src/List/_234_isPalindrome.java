package List;

import DataStructure.ListNode;

import javax.imageio.plugins.tiff.FaxTIFFTagSet;
import java.util.Stack;

public class _234_isPalindrome {

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode p1 = new ListNode(1);
        ListNode p2 = new ListNode(2);
        ListNode p3 = new ListNode(3);
        ListNode p4 = new ListNode(4);
        ListNode p5 = new ListNode(4);
        ListNode p6 = new ListNode(3);
        ListNode p7 = new ListNode(1);
        ListNode p8 = new ListNode(1);
        ListNode p9 = new ListNode(9);

        p1.next=p2;
        p2.next=p3;
        p3.next=p4;
        p4.next=p5;
        p5.next=p6;
        p6.next=p7;
p7.next = p8;
        boolean palindrome = new _234_isPalindrome().isPalindrome(p9);

        if (palindrome) {
            System.out.println("yes The listNode is palindrome ");
        }else {
            System.out.println("No fuck you");
        }

    }
    public boolean isPalindrome(ListNode head) {
        if (head==null) {
            return false;
        }
/*
        int cnt=0;
        ListNode post=head,rev=head;
        while (rev.next != null) {
            cnt++;
            rev=rev.next;
        }*/

        Stack<ListNode> stack = new Stack<>();
        ListNode p1=head,p2=head,cnt=head;
        int numbers =0;
        while (cnt != null) {
            numbers++;
            cnt= cnt.next;
        }

        int mid = numbers / 2;
        while (mid> 0) {
            mid--;
            stack.add(p1);
            p1=p1.next;
            p2=p2.next;
        }
        if (numbers%2==1) p2=p2.next;
        while (p2 != null) {
            ListNode pop = stack.pop();
            if (pop.val!=p2.val) {
                return false;
            }
            p2=p2.next;
        }

        return true;



    }


}
