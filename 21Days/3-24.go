package days

import (
	"container/heap"
)

// 链表排序
type heapLists []*ListNode

func (h heapLists) Len() int {
	return len(h)
}
func (h heapLists) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h heapLists) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *heapLists) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}
func (h *heapLists) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	heaplists := &heapLists{}
	heap.Init(heaplists)
	for head != nil {
		heap.Push(heaplists, head)
		head = head.Next
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for heaplists.Len() != 0 {
		node := heap.Pop(heaplists).(*ListNode)
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = nil
	head = dummyHead.Next
	return head
}

// 两数之和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	n1, n2, carry, cur := 0, 0, 0, dummyHead

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		cur.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		cur = cur.Next
		carry = (n1 + n2 + carry) / 10
	}
	return dummyHead.Next
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root==nil {
		return  nil
	}
	stack,cur:=make([]*TreeNode,0),root
	res:=[]int{}
	for len(stack)!=0||cur!=nil {
		for cur!=nil {
			res=append(res, cur.Val)
			stack=append(stack, cur)
			cur=cur.Left
		}
		node:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		cur=node.Right
	}
	return  res
}