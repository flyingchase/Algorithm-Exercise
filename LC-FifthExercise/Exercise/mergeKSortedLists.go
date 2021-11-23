package exercise

import "container/heap"

func mergeKSortedLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid := length >> 1
	left := mergeKSortedLists(lists[:mid])
	right := mergeKSortedLists(lists[mid:])
	return mergeTwoLists(left, right)
}
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
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

type listHeap []*ListNode

func (l listHeap) Len() int {
	return len(l)
}
func (l listHeap) Less(i int, j int) bool {
	return l[i].Val < l[j].Val
}
func (l listHeap) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *listHeap) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

func (l *listHeap) Pop() interface{} {
	old := *l
	x := old[len(old)-1]
	*l = old[:len(old)-1]
	return x
}

func mergeKSortedLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	listnodeH := &listHeap{}
	heap.Init(listnodeH)
	for _, list := range lists {
		heap.Push(listnodeH, list)
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for listnodeH.Len() != 0 {
		node := heap.Pop(listnodeH).(*ListNode)
		if node.Next != nil {
			heap.Push(listnodeH, node.Next)
		}
		cur.Next = node
		cur, node = cur.Next, node.Next
	}
	return dummyHead.Next
}
