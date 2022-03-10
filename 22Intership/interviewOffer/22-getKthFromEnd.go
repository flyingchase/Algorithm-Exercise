package interviewoffer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && k == 1 {
		return head
	}
	p1, p2 := head, head
	for p1 != nil {
		for k > 0 {
			k--
			p1 = p1.Next
		}
		p1, p2 = p1.Next, p2.Next
	}
	return p2
}
