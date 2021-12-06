package exercise

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	prev, cur := dummyHead, head
	prev.Next = cur
	for i := 0; i < left-1; i++ {
		prev = cur
		cur = cur.Next
	}
	next := cur.Next
	for i := left; i < right; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummyHead.Next
}
