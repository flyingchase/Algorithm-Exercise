package topkLists

// K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 	dummyHead := &ListNode{Val: -1, Next: head}
	// 	cur := head
	// 	count := 0
	// 	for cur != nil {
	// 		count++
	// 		cur = cur.Next
	//
	// 	}
	// 	return dummyHead.Next

	if k == 1 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur != nil {
		end := cur
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummyHead.Next
			}
		}
		prev, next := cur, end.Next
		for x := cur.Next; x != next; {
			prev, x, x.Next = x, x.Next, prev
		}
		cur, cur.Next, cur.Next.Next = cur.Next, end, next
	}
	return dummyHead.Next
}

// 区间逆转链表
func changeBetween(head *ListNode, l, r int) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head
	prev := dummyHead
	cur := head
	for i := 0; i < l-1; i++ {
		prev = cur
		cur = cur.Next
	}
	next := cur.Next
	for i := l; i < r; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummyHead.Next
}
