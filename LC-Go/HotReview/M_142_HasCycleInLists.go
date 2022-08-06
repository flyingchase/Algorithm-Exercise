package HotReview

func hasCycle(head *ListNode) (bool, *ListNode) {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, fast
		}
	}
	return false, head
}

func detectCycle(head *ListNode) *ListNode {
	flag, point := hasCycle(head)
	if !flag {
		return nil
	}
	p1 := head
	for p1 != point {
		p1 = p1.Next
		point = point.Next
	}
	return p1
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle2(head *ListNode) (bool, *ListNode) {
	if head == nil || head.Next == nil {
		return false, nil
	}
	fast, slow := head, head
	flag := false
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			flag = true
			break
		}
	}
	if flag {
		fast = head
		for fast != slow {
			fast, slow = fast.Next, slow.Next
		}
		return true, fast

	}
	return false, nil
}
