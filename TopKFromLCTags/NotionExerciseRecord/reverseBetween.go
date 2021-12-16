package notionexerciserecord

import (
	"topkTags"
)

type ListNode = topkTags.ListNode

func reverseBetween(root *ListNode, left, right int) *ListNode {
	if root == nil || right <= left {
		return root
	}

	dummyHead, cur := &ListNode{Val: -1, Next: root}, root
	prev := dummyHead
	for i := 0; i < left-1; i++ {
		prev = prev.Next
		cur = cur.Next
	}
	next := cur.Next

	for i := left; i < right; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = cur
		next = cur.Next
	}
	return dummyHead.Next
}
