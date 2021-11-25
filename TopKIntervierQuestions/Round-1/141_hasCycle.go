package round1

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// use set (implmented by struct{})
func hasCycle2(head *ListNode) bool {
	dict := make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := dict[cur]; ok {
			return true
		}
		dict[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}
