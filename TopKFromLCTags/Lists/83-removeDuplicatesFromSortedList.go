package topkLists

// iterator
func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	prev.Next = cur
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		prev = cur
		cur = cur.Next
	}
	return dummyHead.Next
}

// recursive
func deleteDuplicate2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicate2(head.Next)
		return head
	} else {
		head.Next = deleteDuplicate2(head.Next)
		return head
	}

}
