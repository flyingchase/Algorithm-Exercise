package HotReview

/*
删除有序链表中的重复数字

*/

// type ListNode = DataStructure.ListNode

func zdeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	// 循环条件为判断直到不相等
	for cur != nil && cur.Next != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		// 每次都下移  非 if
		cur = cur.Next
	}
	return head
}
