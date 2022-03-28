package days

import (
	"container/heap"
)

type maxHeap []int
type minHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i int, j int) bool {
	return m[i] >= m[j]
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

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i int, j int) bool {
	return m[i] <= m[j]
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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	min, max := &minHeap{}, &maxHeap{}
	heap.Init(min)
	heap.Init(max)
	nums1 = append(nums1, nums2...)
	for i := 0; i < len(nums1); i++ {
		if min.Len() == max.Len() {
			heap.Push(max, nums1[i])
			heap.Push(min, heap.Pop(max).(int))
		} else {
			heap.Push(min, nums1[i])
			heap.Push(max, heap.Pop(min).(int))
		}
	}
	var res float64
	if min.Len() == max.Len() {
		res = float64((heap.Pop(max).(int)*1.0 + heap.Pop(min).(int)*1.0)) * 0.5
	} else {
		res = float64(heap.Pop(min).(int))
	}
	return res
}

// 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	// 单调栈正向板子，window 存储 index 下标
	res, window := []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		// 双端队列，保证队首最大
		// shrink
		for len(window) > 0 && window[0] < i-k+1 {
			window = window[1:]
		}
		// 单调栈，保证入栈递增
		// adjust
		for len(window) > 0 && nums[window[len(window)-1]] <= nums[i] {
			window = window[:len(window)-1]
		}
		// extend
		window = append(window, i)
		// ready to in
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	index := len(nums) - 2
	// 第一个降序 num
	for ; index >= 0 && nums[index] > nums[index+1]; index-- {
	}
	if index < 0 {
		// 不存在更大排列
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}
	// 找第一个比 nums[index] 大的数
	j := index + 1
	for ; j < len(nums) && nums[j] > nums[index]; j++ {
	}
	// 多进一位
	j--
	// 替换
	nums[index], nums[j] = nums[j], nums[index]
	// 将 index 之后逆转
	index++
	for j := len(nums) - 1; index < j; j, index = j-1, index+1 {
		nums[index], nums[j] = nums[j], nums[index]
	}
}
