package didi

func findNthNodeFromEndInListNode(root *ListNode, n int) *ListNode {
	if root == nil {
		return nil
	}
	p1, p2, dummyHead := root, root, &ListNode{Val: -1, Next: root}
	prev := p2
	for p1 != nil {
		if n <= 0 {
			prev, p2 = p2, p2.Next
		}
		p1 = p1.Next
		n--
	}
	prev.Next = p2.Next
	return dummyHead.Next
}
