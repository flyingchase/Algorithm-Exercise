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
	for i := 0; i < len(nums1); i++ {
		heap.Push(maxHeap, nums1[i])
	}
	for j := 0; j < len(nums2); j++ {
		heap.Push(minHeap, nums2[j])
	}
}
