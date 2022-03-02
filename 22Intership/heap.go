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

// 合并 k 个有序链表
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

//  最少的不同元素
// nums 删除 k 个元素后剩下不同的元素最少

// 重排字符串，使相同字符不相邻

// 有序矩阵中的第 k 小元素

// 最小范围，包含每个数组的至少一个数

// 查找和最小的 k 对数字
