package exercise

import "container/heap"

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {

		return l2
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	n1, n2, carry := 0, 0, 0
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
		carry = (n1 + n2 + carry) / 10
		cur = cur.Next
	}
	return dummyHead.Next
}
func addTwoSortedArrays3(nums1, nums2 []int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[len(nums2)/2])
	}
	if len(nums2) == 0 {
		return float64(nums1[len(nums1)/2])
	}
	nums1 = append(nums1, nums2...)
	max, min := &maxHeap{}, &minHeap{}
	heap.Init(max)
	heap.Init(min)
	for _, num := range nums1 {
		if max.Len() == min.Len() {
			heap.Push(max, num)
			heap.Push(min, heap.Pop(max))
		} else {
			heap.Push(min, num)
			heap.Push(max, heap.Pop(min))
		}
	}
	return findMedian(max, min)
}

func findMedian(max *maxHeap, min *minHeap) float64 {
	if max.Len() == min.Len() {
		return float64((heap.Pop(max).(int) + heap.Pop(min).(int)) >> 1)
	}
	return float64(heap.Pop(min).(int))
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i int, j int) bool {
	return m[i] > m[j]
}
func (m maxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i int, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '-' {
			sign = -1
		}
	}
	return sign * res
}
