package reviewlc60

import "alg"

type ListNode = alg.ListNode

func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	for cur != nil {
		next := cur.Next
		if cur.Val == next.Val {
			prev.Next = next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
