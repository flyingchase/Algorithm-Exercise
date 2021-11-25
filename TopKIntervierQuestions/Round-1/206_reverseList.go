package round1

import (
	"topk"
)

type ListNode = topk.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur:=head
	// &ListNode{}赋值默认为零值，导致最后多一位 0
	// var *ListNode 默认为 空，
	var prev *ListNode
	for cur != nil {
		next:= cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
