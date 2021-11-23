package List;

import InnerStruct.ListNode;

/* 
- 先利用快慢指针。若能相遇，说明存在环，且相遇点一定是在环上；若没有相遇，说明不存在环，返回 `null`。
- 固定当前相遇点，用一个指针继续走，同时累积结点数。计算出环的结点个数 `cnt`。
- 指针 p1 先走 `cnt` 步，p2 指向链表头部，之后 `p1`,`p2` 同时走，相遇时，相遇点一定是在环的入口处。因为 `p1` 比 `p2` 多走了环的一圈。
*/
public class EntryNodeOfLoop {
    // 快慢双指针法
    public ListNode ntryNodeOfLoop(ListNode pHeap) {
        if (pHeap == null || pHeap.next == null) {
            return null;
        }
        ListNode slow = pHeap, fast = pHeap;
        boolean flag = false;
        // 确保有环
        while (fast.next != null && fast != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (fast == slow) {
                flag = true;
                break;
            }
        }
        // 链表无环
        if (!flag) {
            return null;
        }
        // 记录下快慢指针相遇点
        ListNode cur = slow.next;
        // 求出环的长度
        int cnt = 1;
        while (cur != slow) {
            cur = cur.next;
            cnt++;
        }

        // 指针p1从头节点开始走cnt环长
        ListNode p1 = pHeap;
        for (int i = 0; i < cnt; i++) {
            p1 = p1.next;
        }

        // 指针p2从头节点开始与p1同步走 首次相遇即为环入口
        ListNode p2 = pHeap;
        while (p1 != p2) {
            p1 = p1.next;
            p2 = p2.next;
        }
        return p1;
    }
}
