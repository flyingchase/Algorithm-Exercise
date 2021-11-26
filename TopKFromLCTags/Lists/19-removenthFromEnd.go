package topkLists

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev := dummyHead
	fast, slow := head, head
	for fast != nil {
		for n > 0 {
			n--
			fast = fast.Next
		}
		if fast == nil {
			break
		}
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next
	return dummyHead.Next
}
