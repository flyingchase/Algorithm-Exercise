package datawhale50

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	n1, n2, carry, cur := 0, 0, 0, dummyHead
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		cur.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		cur = cur.Next
		carry = (n1 + n2 + carry) / 10
	}
	return dummyHead.Next
}
