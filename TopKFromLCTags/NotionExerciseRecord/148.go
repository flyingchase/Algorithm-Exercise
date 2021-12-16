package notionexerciserecord

import (
	"container/heap"
)

type PQ []*ListNode

func (p PQ) Len() int {
	return len(p)
}
func (p PQ) Less(i int, j int) bool {
	return p[i].Val < p[j].Val
}

func (p PQ) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*ListNode))
}
func (p *PQ) Pop() interface{} {
	old := *p
	x := old[:len(old)-1]
	*p = old[:len(old)-1]
	return x
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pq := &PQ{}
	heap.Init(pq)
	dummyHead, cur := new(ListNode), head
	dummyHead.Next = head
	for cur != nil {
		heap.Push(pq, cur)
		cur = cur.Next
	}
	cur = dummyHead
	for pq.Len() >= 0 {
		node := heap.Pop(pq).(*ListNode)
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}
