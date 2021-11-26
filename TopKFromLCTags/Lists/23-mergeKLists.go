package topkLists

import (
	"container/heap"
)

type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i int, j int) bool {
	return m[i].Val < m[j].Val
}
func (m minHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	m := &minHeap{}
	heap.Init(m)
	for i := 0; i < len(lists); i++ {
		// attention 当 lists 内含有空链表时，不入堆
		if lists[i] != nil {
			heap.Push(m, lists[i])
		}
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for m.Len() > 0 {
		node := heap.Pop(m).(*ListNode)
		if node.Next != nil {
			heap.Push(m, node.Next)
		}
		cur.Next = node
		cur, node = cur.Next, node.Next
	}
	return dummyHead.Next
}
