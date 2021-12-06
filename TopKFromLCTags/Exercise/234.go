package exercise

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var dummyHead *ListNode

	slow, fast, prev := head, head, dummyHead
	prev.Next = head
	for fast != nil && fast.Next != nil {
		slow.Next = prev
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}
