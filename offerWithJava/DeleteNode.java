

import Function.ListOperate;
import InnerStruct.ListNode;

public class DeleteNode {
    public static ListNode deleteNode(ListNode head, ListNode tobeDelete) {
        if (head == null || tobeDelete == null) {
            return null;
        }
        // 非尾节点删除
        if (tobeDelete != null) {
            ListNode next = tobeDelete.next;
            tobeDelete.val = next.val;
            tobeDelete.next = next.next;
        } else {
            if (head == tobeDelete) {
                // 仅仅只有尾节点
                head = null;
            } else {
                ListNode cur = head;
                while (cur.next != tobeDelete) {
                    cur = cur.next;
                }
                cur.next = null;
            }
        }
        return head;
    }

    public static void main(String[] args) {
        ListOperate listOperate = new ListOperate();
        int[] nums = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 };
        ListNode head = listOperate.arrayToListNode(nums);
        ListNode tobeDelete = new ListNode(1,null);
        head=deleteNode(head, tobeDelete);
        listOperate.printListNode(head);
    }
}
