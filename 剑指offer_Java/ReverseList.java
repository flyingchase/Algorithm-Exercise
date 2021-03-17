package 剑指offer_Java;

import java.util.ArrayList;
import java.util.Stack;
import 剑指offer_java.*;
/*
输入一个链表，按链表值从尾到头的顺序返回一个 ArrayList。
*/


import javax.swing.ListCellRenderer;

public class ReverseList {
   

    public class ListNode {
        int val;
        ListNode next = null;

        public ListNode(int val) {
            this.val = val;
        }
    }
    public ArrayList<Integer> reverseList(ListNode listNode){
        ArrayList<Integer> res = new ArrayList<>();
        if(listNode== null) return res;
        Stack<Integer> stack = new Stack<>();
        while (listNode!=null) {
            stack.push(listNode.val);
            listNode = listNode.next;
            
        }
        while(!stack.isEmpty()){
            res.add(stack.pop());
        }
        return res;
    }

    public static void main(String[] args) {
        int[] s= {1,2,3,4,5,6,7,8,9,0};
        ReverseList r1= new ReverseList();
        
        ListNode root = new ReverseList().new ListNode(s[0]);
        ListNode nextNode = root;
        for(int i=1;i<s.length;i++){
            
            ListNode tempNode=new ReverseList().new ListNode(s[i]);
            nextNode.next=tempNode;
            nextNode=tempNode;
        }
        ReverseList reslist = new ReverseList();
        ArrayList<Integer> res = reslist.reverseList(root);

        for(Integer m : res) System.out.print(m+" ");
        System.out.println();
    }

}

