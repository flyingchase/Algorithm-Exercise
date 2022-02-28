package intership

// 判断链表是否有环
// 快慢指针，起步快指针指向 head，慢指针指向 head.Next
// 一旦 fast 到达末尾则无环，否则最后一定有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head.Next
	for slow != fast {
		if fast.Next == nil || fast == nil {
			return false
		}
		fast, slow = fast.Next.Next, slow.Next
	}
	return true
}

// 判断链表是否有环
// 使用 dict 字典，其中空结构体为 value 占位
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]struct{})
}
func detechCycle(head *ListNode) (*ListNode, bool) {
	if head == nil || head.Next == nil {
		return head, false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return fast, true

		}
	}
	return head, false
}
