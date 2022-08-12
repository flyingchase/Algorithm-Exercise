package HotReview

/*区间链表逆转*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}

	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead

	for i := 0; i < m-1; i++ {
		cur = cur.Next
	}
	wall := cur
	leftNode := wall.Next
	rightNode := leftNode.Next

	for i := m + 1; i <= n; i++ {
		leftNode.Next = rightNode.Next
		rightNode.Next = wall.Next
		wall.Next = rightNode
		// rightN 后移
		rightNode = leftNode.Next
	}
	return dummyHead.Next
}
