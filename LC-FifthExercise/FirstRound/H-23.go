package firstround

import (
	"container/heap"
)

type ListNodeList []*ListNode
type ListNode struct {
	Next *ListNode
	Val  int
}

func (h ListNodeList) Len() int {
	return len(h)
}
func (h ListNodeList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h ListNodeList) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *ListNodeList) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeList) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mergeKSortedLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	h := &ListNodeList{}
	heap.Init(h)
	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}
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
