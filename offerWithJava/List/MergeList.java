package List;

import Function.ListOperate;
import InnerStruct.ListNode;
import jdk.internal.jshell.tool.resources.l10n;

/* 
合并有序链表
 */
public class MergeList {
    // 迭代
    public static ListNode Merge(ListNode list1, ListNode list2) {
        if (list1==null) {
            return list2;
        }
        else if(list2==null){
            return list1;
        }
        ListNode p1=list1,p2=list2;
        ListNode head = new ListNode(-1);
        ListNode cur=head;
        while (list1!=null&&list2!=null) {
            if(list1.val<list2.val) {
                cur.next=list1;
                list1=list1.next;
            }
            else{
                cur.next=list2;
                list2=list2.next;
            }
            cur=cur.next;
        }
        // 不同于数组 有未合并的部分直接将节点纳入辅助节点后即可 无需遍历
        if(list1!=null) {
            cur.next=list1;
        }
        if(list2!=null) {
            cur.next=list2;
        }
        return head.next;
    }    

    // 递归
    public static ListNode Merge02(ListNode list1, ListNode list2) {
        if(list1==null) {
            return list2;
        }
        if(list2==null) {
            return list1;
        }
        if(list1.val<list2.val) {
            list1.next=Merge02(list1.next, list2);
            return list1;
        }
        list2.next=Merge02(list1, list2.next);
        return list2;
        
    }

    public static void main(String[] args) {
        int[] nums1={1,3,5,7,9};
        int[] nums2 = {2,4,6,8,10};
        ListOperate listOperate = new ListOperate();
        ListNode list1= listOperate.arrayToListNode(nums1);
        ListNode list2= listOperate.arrayToListNode(nums2);
        ListNode head = Merge(list1, list2);
        listOperate.printListNode(head);
        
        list1= listOperate.arrayToListNode(nums1);
        list2= listOperate.arrayToListNode(nums2);
        ListNode head2=Merge02(list1, list2);
        listOperate.printListNode(head2);
    }
}
