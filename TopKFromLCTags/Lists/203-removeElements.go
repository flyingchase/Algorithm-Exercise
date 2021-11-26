package topkLists

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
			continue
		}
		prev = prev.Next
	}
	return dummyHead.Next
}
