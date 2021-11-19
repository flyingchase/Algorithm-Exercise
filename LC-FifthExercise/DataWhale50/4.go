package datawhale50

import (
	"container/heap"
)

// 构造大小堆分别存储数组中较大的一部分和较小的一部分
type maxHeap []int
type minHeap []int

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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
	maxH, minH := &maxHeap{}, &minHeap{}
	// 初始化大小堆
	heap.Init(maxH)
	heap.Init(minH)
	nums1 = append(nums1, nums2...)
	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		// 大小堆 size 相同则先入大堆，再将大堆堆顶弹出入小堆
		// 保证小堆存储较大部分
		if maxH.Len() == minH.Len() {
			heap.Push(maxH, num)
			heap.Push(minH, heap.Pop(maxH))
		} else {
			// 大小堆 size 不等则先入小堆，再小堆堆顶弹出入大堆
			// 保证大堆存储较小部分
			heap.Push(minH, num)
			heap.Push(maxH, heap.Pop(minH))
		}
	}
	return findMedian(maxH, minH)
}

// 大小堆相等则取两堆堆顶的平均数
// 否则取小堆堆顶为中位数
func findMedian(maxH *maxHeap, minH *minHeap) float64 {
	if len(*maxH) == len(*minH) {
		return float64((heap.Pop(maxH).(int)+heap.Pop(minH).(int))*1.0) * 0.5
	} else {
		return (float64(heap.Pop(minH).(int) * 1.0))
	}
}
