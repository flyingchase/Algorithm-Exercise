 package Function;
import InnerStruct.*;

public class ListOperate {
    // 将Integer[] nums转为为链表
    public ListNode arrayToListNode(int[] nums) {
        if(nums.length==0) return null;
        ListNode head = new ListNode(nums[0]);
        ListNode other = head;
        for (int i = 1; i < nums.length; i++) {
            ListNode temp = new ListNode(nums[i]);
            other.next=temp;
            other=temp;
        }
        return head;
    }
    public  void printListNode(ListNode head) {
        ListNode temp =head;
        while (temp!=null) {
            System.out.print(temp.val+" ");
            temp=temp.next;
        }
        System.out.println();
    }
}
