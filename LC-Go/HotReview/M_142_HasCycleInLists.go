package HotReview

func hasCycle(head *ListNode) (bool, *ListNode) {

	fast,slow:=head,head
	for fast!=nil&&fast.Next!=nil {
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow {
			return true,fast
		}
	}
	return false,head
}

func detectCycle(head *ListNode) *ListNode {
	flag,point:=hasCycle(head)
	if !flag {

	return nil
	}
	p1:=head
	for p1!=point {
		p1=p1.Next
		point= point.Next
	}
	return p1
}
