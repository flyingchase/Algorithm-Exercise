package exercise

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
func hasCycle2(head *ListNode) bool {
	dict := make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		dict[cur] = struct{}{}
		if _, ok := dict[cur]; ok {
			return true
		}
		cur = cur.Next
	}
	return false
}
