package datawhale50

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	cur1, cur2 := l1, l2
	cur := dummyHead
	others := 0
	for cur1 != nil && cur2 != nil || others != 0 {
		if cur1 != nil && cur2 != nil {
			others += cur1.Val + cur2.Val
		}
		cur.Next = &ListNode{Val: others % 10}
		cur = cur.Next
		others /= 10
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	return dummyHead.Next
}
