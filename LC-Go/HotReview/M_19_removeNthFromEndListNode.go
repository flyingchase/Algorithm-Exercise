package HotReview

func removerNthFromEndListNode(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, p1, p2 := dummyHead, head, head
	for p2 != nil {
		if n <= 0 {
			prev = p1
			p1 = p1.Next
		}
		n--
		p2 = p2.Next
	}
	prev.Next = p1.Next
	return dummyHead.Next
}

func removeNthFromEndListNode1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head

}

func removerNthFromEndListNode2(root *ListNode, n int) *ListNode {
	if root == nil {
		return nil
	}
	dummyHead := &ListNode{Next: root}
	prev, p1, p2 := dummyHead, root, root

	for p1 != nil {
		p1 = p1.Next
		n--
		if n < 0 {
			prev = prev.Next
			p2 = p2.Next
		}
	}
	prev.Next = p1.Next
	return dummyHead.Next
}
