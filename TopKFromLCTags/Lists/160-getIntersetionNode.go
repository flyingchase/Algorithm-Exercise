package topkLists

import "math"

// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	n1, n2 := 0, 0
	for l1 != nil {
		n1++
		l1 = l1.Next
	}
	for l2 != nil {
		n2++
		l2 = l2.Next
	}
	gap := int(math.Abs(float64(n1 - n2)))
	l1, l2 = headA, headB
	if n1 > n2 {
		for gap > 0 {
			l1 = l1.Next
			gap--
		}
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next
	} else {
		for gap > 0 {
			gap--
			l2 = l2.Next
		}
		if l1 == l2 {
			return l2
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return nil
}
