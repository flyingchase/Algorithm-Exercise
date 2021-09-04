package HotReview

/*给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	//DataStructure.ListNode{}
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)

	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != nil {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp

	}
	return prev

}
