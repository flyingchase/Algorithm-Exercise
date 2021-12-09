package exercise

import "container/heap"

type minHeap []int
type maxHeap []int

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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) == 0 {
		return float64(nums1[len(nums1)>>1])
	}
	if len(nums2) == 0 {
		return float64(nums2[len(nums2)>>1])
	}
	nums1 = append(nums1, nums2...)
	max, min := &maxHeap{}, &minHeap{}
	heap.Init(max)
	heap.Init(min)
	for _, num := range nums1 {
		if max.Len() == min.Len() {
			// 两堆相等则先入大堆
			// 此处错误
			heap.Push(min, num)
			heap.Push(max, heap.Pop(min))
		} else {
			heap.Push(max, num)
			heap.Push(min, heap.Pop(max))
		}
	}
	var res float64
	if max.Len() == min.Len() {
		res = float64(heap.Pop(max).(int)+heap.Pop(min).(int)) * 0.5
	} else {
		res = float64(heap.Pop(min).(int))
	}

	return res
}
