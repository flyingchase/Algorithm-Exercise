package datawhale50

import "container/heap"

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size < 1 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}

	// recursive
	mid := size >> 1
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}
func mergeKLists2(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	h := &listHeap{}
	heap.Init(h)
	// 各表头入堆
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	// dfs
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for h.Len() != 0 {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		cur.Next = node
		cur, node = cur.Next, node.Next
	}
	return dummyHead.Next
}
