package sorts

import (
	"alg"
	"container/heap"
)

type ListNode = alg.ListNode

// H-23-合并 k 个有序链表
// soulution_1: 不断分割后合并 类似 mergesort
func mergeKLists_1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeKListsHelper(lists, 0, len(lists)-1)
}
func mergeKListsHelper(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l < r {
		mid := l + (r-l)>>1
		l1 := mergeKListsHelper(lists, l, mid)
		l2 := mergeKListsHelper(lists, mid+1, r)
		return mergeLists(l1, l2)
	}
	return nil
}
func mergeLists(l1, l2 *ListNode) *ListNode {
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

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// 使用堆实现合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	minHeap := &minHeap{}
	heap.Init(minHeap)
	for _, list := range lists {
		// 保证非空情况才入堆
		if list != nil {
			heap.Push(minHeap, list)
		}
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for minHeap.Len() != 0 {
		node := heap.Pop(minHeap).(*ListNode)
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
		cur.Next = node
		cur, node = cur.Next, node.Next
	}
	return dummyHead.Next
}
