package List;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Stack;

import Function.*;
import InnerStruct.*;
/**
 * PrintListInReversedOrder 
 * 输入一个链表，按链表值从尾到头的顺序返回一个 `ArrayList`。
 * 
 * 
 * 
 */


public class PrintListInReversedOrder {
    
    // 递归 
    public static ArrayList<Integer> print01ReversedOrderListNode(ListNode listNode) {
        ArrayList<Integer> ret = new ArrayList<>();
        if(listNode != null) {
            ret.addAll(print01ReversedOrderListNode(listNode.next));
            ret.add(listNode.val); 
        }
        return ret;
    }
    // 头插法
    public static ArrayList<Integer> print02ReversedOrderListNode(ListNode listNode) {
        ListNode head = new ListNode(-1); // 非真头节点 用来统一插入操作        
        while (listNode!=null) {
            ListNode memo = listNode.next;
            listNode.next=head.next;
            head.next=listNode;
            listNode=memo;
        }
        ArrayList<Integer> ret = new ArrayList<>();
        head =head.next;
        while(head!=null){
            ret.add(head.val);
            head=head.next;
        }
        return ret;
    }

    // 辅助栈
    public static ArrayList<Integer> print03ReversedOrderListNode(ListNode listNode) {
        Stack<Integer> stack = new Stack<>();
        while (listNode!=null) {
            stack.add(listNode.val);
            listNode=listNode.next;
        }
        ArrayList<Integer> ret = new ArrayList<>();
        while (!stack.isEmpty()) {
            ret.add(stack.pop());
        }
        return ret;
    }
    public static void main(String[] args) {
        int[] nums={1,2,3,4,5,6,7,8,9,0};
        Integer[] integers = Arrays.stream(nums).boxed().toArray(Integer[]::new);
        ListOperate listOperate = new ListOperate();

        ListNode head = ((ListOperate) listOperate).arrayToListNode(nums);
        
        ArrayList<Integer> arr = new ArrayList<>();
        arr= print01ReversedOrderListNode(head);
        System.out.println(arr.toString());
       
        arr= print03ReversedOrderListNode(head);
        System.out.println(arr.toString());
    }
}