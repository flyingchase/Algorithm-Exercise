package topkLists

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: list1}
	prev1, cur := dummyHead, list1
	prev1.Next = cur
	prev2 := dummyHead
	prev2.Next = cur
	for i := 0; i < a; i++ {
		prev1 = cur
		prev2 = cur
		cur = cur.Next
	}
	prev1.Next = list2
	cur2 := list2
	for cur2.Next != nil {
		cur2 = cur2.Next
	}
	for i := 0; i < b-a+1; i++ {
		prev2 = cur
		cur = cur.Next
	}
	cur2.Next = cur
	return dummyHead.Next
}
