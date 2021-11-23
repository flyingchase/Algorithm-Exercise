package topselection

import "container/heap"

type intervals []int

func (i intervals) Len() int {
	return len(i)
}
func (h intervals) Less(i int, j int) bool {
	return h[i] > h[j]
}
func (h intervals) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (i *intervals) Push(x interface{}) {
	*i = append(*i, x.(int))
}

func (i *intervals) Pop() interface{} {
	old := *i
	x := old[len(old)-1]
	*i = old[:len(old)-1]
	return x
}
func findKthLargest(nums []int, k int) int {

	if len(nums) == 0 || k > len(nums) {
		return 0
	}
	h := &intervals{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	for k > 1 {
		heap.Pop(h)
		k--
	}
	return heap.Pop(h).(int)
}
