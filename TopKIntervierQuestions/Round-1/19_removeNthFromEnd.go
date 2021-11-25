package round1

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	p1, p2 := head, head
	prev := dummyHead
	for p1 != nil {
		for n > 0 {
			n--
			p1 = p1.Next
		}
		if p1 == nil {
			break
		}
		prev = p2
		p2 = p2.Next
		p1 = p1.Next
	}
	prev.Next = p2.Next
	return dummyHead.Next
}
