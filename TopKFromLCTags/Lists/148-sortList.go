package topkLists
//
// import "container/heap"
//
// type minHeap []*ListNode
//
// func (m minHeap) Len() int {
// 	return len(m)
// }
// func (m minHeap) Less(i int, j int) bool {
// 	return m[i].Val < m[j].Val
// }
// func (m minHeap) Swap(i int, j int) {
// 	m[i], m[j] = m[j], m[i]
// }
// func (m *minHeap) Push(x interface{}) {
// 	*m = append(*m, x.(*ListNode))
// }
// func (m *minHeap) Pop() interface{} {
// 	old := *m
// 	x := old[len(old)-1]
// 	*m = old[:len(old)-1]
// 	return x
// }
//
// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	m := &minHeap{}
// 	heap.Init(m)
// 	for head != nil {
// 		heap.Push(m, head)
// 		head = head.Next
// 	}
// 	dummyHead := &ListNode{Val: -1}
// 	cur := dummyHead
// 	for m.Len() > 0 {
// 		cur.Next = heap.Pop(m).(*ListNode)
// 		cur = cur.Next
// 	}
// 	return dummyHead.Next
// }


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	prev:=head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 利用快慢指针找到中点
	// slow即为中点，断开 slow 的前结点链接
	prev.Next = nil
	return  merge(sortList(head),sortList(slow))
}
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead:=new(ListNode)
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l1.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummyHead.Next
}

/*
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	l1 := sortList(head)
	l2 := sortList(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := new(ListNode)
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}
*/