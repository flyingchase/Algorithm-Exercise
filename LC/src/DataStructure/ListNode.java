package DataStructure;



public class ListNode {
    public int val;
    public ListNode Next;

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.Next = next;
    }
}