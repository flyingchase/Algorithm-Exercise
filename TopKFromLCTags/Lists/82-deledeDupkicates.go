package topkLists

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	prev.Next = cur
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
