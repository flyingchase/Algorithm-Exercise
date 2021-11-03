package secondround

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteNthFromEndListNode(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}
	dummyHead, slow, fast := &ListNode{Next: head}, head, head
	prevSlow := dummyHead
	for fast != nil {
		if n <= 0 {
			prevSlow, slow = slow, slow.Next
		}
		fast = fast.Next
		n--
	}
	prevSlow.Next = slow.Next
	return dummyHead.Next
}
