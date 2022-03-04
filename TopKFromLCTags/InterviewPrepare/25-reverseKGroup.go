package interviewprepare

type ListNode struct {
	Next *ListNode
	Val  int
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	cur = head
	i := 1
	for ; i <= length-k; i += k {
		cur = prev.Next
		prev.Next = reverseBetweenLAndR(cur, i, i+k-1)
	}
	if (length-i+1)&k == 0 {
		cur = prev.Next
		prev.Next = reverseBetweenLAndR(cur, length-i, length)

	}
	return dummyHead.Next
}
func reverseBetweenLAndR(root *ListNode, l, r int) *ListNode {
	if root == nil || l >= r {
		return root
	}
	dummyHead, cur := &ListNode{Val: -1, Next: root}, root
	prev := dummyHead
	for i := 0; i < l-1; i++ {
		prev = cur
		cur = cur.Next
	}
	next := cur.Next
	for i := l; i < r; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummyHead.Next
}
