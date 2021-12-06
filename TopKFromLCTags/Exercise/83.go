package exercise

import "topkTags"

type ListNode = topkTags.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	prev := dummyHead
	cur := head
	for cur != nil {
		next := cur.Next
		if next.Val == cur.Val {
			prev.Next = next
			cur = next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
