import org.w3c.dom.Node;

import java.util.Map;
@SuppressWarnings({"all"})
public class Cache {
    private Map<Integer, Node> cache;
    public int get(int key) {
        return key;
    }
    public void put(int key, int value) {}
    class DoubleLinkedList{}

}

class LFUNode{
    int key;
    int value;
    int frequency;
    LFUNode prev;
    LFUNode next;

    public LFUNode(int key, int value) {
        this.key = key;
        this.value = value;
        this.frequency = 1;
    }
}

class DoubleLinkedList {
    int lisSize;
    LFUNode head;
    LFUNode tail;

    public DoubleLinkedList() {
        this.lisSize = 0;
        this.head=new LFUNode(0,0);
        this.tail=new LFUNode(0,0);
        head.next=tail;
        tail.prev=head;
    }
    // 双向链表头插法
    public void addNode(LFUNode curNode) {
        LFUNode nextNode = head.next;
        curNode.next=nextNode;
        curNode.prev=head;
        head.next=curNode;
        nextNode.prev=curNode;
        lisSize++;
    }

    //双向链表删除结点
    public void removeNode(LFUNode curNode) {
        LFUNode prevNode=curNode.prev;
        LFUNode nextNode = curNode.next;
        prevNode.next=nextNode;
        nextNode.prev=prevNode;
        lisSize--;
    }

    // 双向链表删除尾结点
    public LFUNode removeTail() {
        if (lisSize>0) {
            LFUNode tailNode=tail.prev;
            removeNode(tailNode);
            return tailNode;
        }
        return null;
    }
}
