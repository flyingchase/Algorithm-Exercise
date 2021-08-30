package List;

import DataStructure.ListNode;

/*请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。*/
public class E_237_deletNode {
    public void deleteNode(ListNode node) {
        // cover the node val with the next
        node.val=node.next.val;
        // then delete the next Node
        node.next=node.next.next;
    }
}
