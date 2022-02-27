package intership

import "container/heap"

// 合并 k 个有序链表
type heapLists []*ListNode

func (h heapLists) Len() int {
	return len(h)
}
func (h heapLists) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapLists) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}
func (h *heapLists) Pop() interface{} {
	old := *h
	x := old[h.Len()-1]
	old = old[:h.Len()-1]
	*h = old
	return x
}
func (h *heapLists) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func mergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	h := &heapLists{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		heap.Push(h, lists[i])
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for h.Len() != 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		cur = node
	}
	return dummyHead.Next
}
