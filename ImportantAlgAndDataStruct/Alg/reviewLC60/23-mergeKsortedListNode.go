package reviewlc60

func mergeKsortedListNodes(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeHelper(lists, 0, len(lists)-1)
}
func mergeHelper(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	mid := l + (r-l)>>1
	left := mergeHelper(lists[:mid], l, mid)
	right := mergeHelper(lists[mid:], mid+1, r)

	return merge(left, right)
}
func merge(l, r *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for l != nil && r != nil {
		if l.Val <= r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return dummyHead.Next
}
