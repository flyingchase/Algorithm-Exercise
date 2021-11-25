package thirdround

import (
	"container/heap"
)

/*
 给定大小的数组 m,n 正序排列
*/
type maxHeap []int
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h maxHeap) Len() int {
	return len(h)
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[len(nums2)/2])
	}
	if len(nums2) == 0 {
		return float64(nums1[len(nums1)/2])
	}
	maxHeap := &maxHeap{}
	minHeap := &minHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	nums1 = append(nums1, nums2...)
	for _, num := range nums1 {
		if maxHeap.Len() == minHeap.Len() {
			heap.Push(maxHeap, num)
			heap.Push(minHeap, heap.Pop(maxHeap))
		} else {
			heap.Push(minHeap, num)
			heap.Push(maxHeap, heap.Pop(minHeap))
		}
	}
	if minHeap.Len() == maxHeap.Len() {
		return float64(heap.Pop(maxHeap).(int)>>1 + heap.Pop(minHeap).(int)>>1)
	}
	return float64(heap.Pop(minHeap).(int))

}
