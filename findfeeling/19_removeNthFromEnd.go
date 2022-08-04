package findfeeling

import "findfeeling/model"

type ListNode = model.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	prev, cur := dummyHead, head
	for cur != nil {
		prev, cur = prev.Next, cur.Next
		n--
		if n == 0 {
			break
		}
	}
	prev.Next = cur.Next
	return dummyHead.Next
}
