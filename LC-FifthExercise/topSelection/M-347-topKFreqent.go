package topselection

import "container/heap"

type frequency struct {
	num   int
	count int
}

type frequenceCount []*frequency

func (f frequenceCount) Len() int {
	return len(f)
}
func (f frequenceCount) Less(i int, j int) bool {
	return f[i].count > f[j].count
}
func (f frequenceCount) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f *frequenceCount) Push(x interface{}) {
	*f = append(*f, x.(*frequency))
}
func (f *frequenceCount) Pop() interface{} {
	old := *f
	x := old[len(old)-1]
	*f = old[:len(old)-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return nil
	}
	h := &frequenceCount{}
	heap.Init(h)
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}
	for num, count := range m {
		f := &frequency{
			num:   num,
			count: count,
		}
		heap.Push(h, f)
	}
	res := make([]int, 0)
	for k > 0 {
		res = append(res, heap.Pop(h).(*frequency).num)
		k--
	}
	return res
}
