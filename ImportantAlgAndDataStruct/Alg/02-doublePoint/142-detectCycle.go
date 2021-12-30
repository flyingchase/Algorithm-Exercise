package doublepoint

import (
	"alg"
)

type ListNode = alg.ListNode

// 142-找到环入口
// 一遍遍历，不用 hasCycle 函数
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return slow
		}
	}
	return nil
}
func hasCycle(head *ListNode) (bool, *ListNode) {
	if head == nil || head.Next == nil {
		return false, nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true, fast
		}
	}
	return false, nil
}
