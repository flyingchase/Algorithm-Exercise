package topkLists

// 区间逆转链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	cur, prev, next := head, dummyHead, head
	// 先走 left-1 步
	for i := 0; i < left-1; i++ {
		cur = cur.Next
		prev = prev.Next
	}
	next = cur.Next
	// 在 left 到 right 之间逐个头插链表, 走right-1-left+1 次
	for i := left; i < right; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummyHead.Next
}
