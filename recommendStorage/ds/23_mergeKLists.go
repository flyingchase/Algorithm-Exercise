package ds

import "container/heap"

// H_23_合并 k 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	size := len(lists)
	mid := size >> 1
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	l1, l2 := list1, list2
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

// heap sort
type listHeaep []*ListNode

func (h listHeaep) Len() int {
	return len(h)
}
func (h listHeaep) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h listHeaep) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}
func (h *listHeaep) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *listHeaep) Push(x interface{}) {
	(*h) = append((*h), x.(*ListNode))
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	nodeHeap := &listHeaep{}
	heap.Init(nodeHeap)
	for _, listnode := range lists {
		if listnode != nil {
			heap.Push(nodeHeap, listnode)
		}
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for nodeHeap.Len() > 0 {
		node := heap.Pop(nodeHeap).(*ListNode)
		if node.Next != nil {
			heap.Push(nodeHeap, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}
