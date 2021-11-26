package topkLists

import (
	"topkTags"
)

type ListNode = topkTags.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
