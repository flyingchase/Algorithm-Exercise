package HotReview

import (
	_ "LC-Go/DataStructure"
)

// attention  循环终止条件 两链表任意一条走到到末尾或者余数为 0
// in case (remain <10 and the two lists come to the nil)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// construct dummyHead to avoid debate on head and the true head can be found
	// by the dummyHead.Next
	dummyHead := &ListNode{Val: -1}
	n1, n2, carry, cur := 0, 0, 0, dummyHead

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		carry = (n1 + n2 + carry) / 10
		cur = cur.Next
	}
	return dummyHead.Next
}
