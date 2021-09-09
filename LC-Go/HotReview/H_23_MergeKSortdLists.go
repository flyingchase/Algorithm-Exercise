package HotReview

import (
	"LC-Go/DataStructure"
	"container/heap"
)

/*
K个有序链表合并
	- bfs 层序遍历+Heap
		- 每个链表的表投入 map
		- 逐个比较 map[index] 只要存在 next 则继续下移入 queue
	- recursive
		- mid merge
*/
type ListNode = DataStructure.ListNode

//type
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil

	}
	if length == 1 {
		return lists[0]

	}

	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeSelf(left, right)
}

func mergeSelf(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeSelf(left.Next, right)
		return left
	}
	right.Next = mergeSelf(right.Next, left)
	return right
}

func mergeTwoListSorted(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := &ListNode{-1, nil}
	cur := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur = l2.Next
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l2 != nil {
		cur.Next = l2
	}
	if l1 != nil {
		cur.Next = l1
	}
	return dummyHead.Next
}

// use heap

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}
func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	// 使用前初始化
	h := &ListNodeHeap{}
	heap.Init(h)
	// 将 lists 中的单链表头全部入堆
	for _, item := range lists {
		if item != nil {
			heap.Push(h, item)
		}
	}

	head := &ListNode{Val: -1}
	cur := head
	// 遍历 heap
	for h.Len() != 0 {
		node := heap.Pop(h).(*ListNode)
		// 堆顶存在后续结点 将其入堆
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		cur.Next = node
		cur, node = cur.Next, node.Next
	}
	return head.Next

}
