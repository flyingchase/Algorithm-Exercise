package topkLists

func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	for cur.Next != nil {
		next := cur.Next
		if next.Val == cur.Val {
		}

	}

}
