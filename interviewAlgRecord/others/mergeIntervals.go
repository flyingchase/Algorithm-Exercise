package others

import (
	"container/heap"
	"math"
	"sort"
)

/*
给定一组数据，每个数据中含有 rank 和 number 两个值，找出 rank 与 number 最接近的数组，
相同最接近时，按照 rank 最小输出
*/

// 即构造数据结构存储，rank number sore=abs(number-rank)
// 按照 sore 排序，sore 相同时候按照 rank 较小排序
// 手写排序算法、或者使用 sort.slice API

type interval struct {
	rank   int
	number int
	sore   int
}

type intervals []interval

func (in intervals) Len() int {
	return len(in)
}
func (in intervals) Less(i int, j int) bool {
	if in[i].sore < in[j].sore {
		return true
	} else if in[i].sore > in[j].sore {
		return false
	}
	return in[i].rank < in[j].rank
}
func (in intervals) Swap(i int, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in *intervals) Push(x interface{}) {
	*in = append(*in, x.(interval))
}

func (in *intervals) Pop() interface{} {
	old := *in
	x := old[len(old)-1]
	*in = old[:len(old)-1]
	return x
}

func mergeIntervals(nums map[int]int) int {
	if len(nums) == 0 {
		return -1
	}

	in := &intervals{}
	heap.Init(in)
	for k, v := range nums {
		i := interval{
			rank:   k,
			number: v,
			sore:   int(math.Abs(float64(k - v))),
		}
		heap.Push(in, i)
	}
	return heap.Pop(in).(interval).rank
}
func mergeIntervals2WithInnerAPI(nums map[int]int) int {
	if len(nums) == 0 {
		return -1
	}
	var ins []interval
	for k, v := range nums {
		in := interval{
			rank:   k,
			number: v,
			sore:   int(math.Abs(float64(k - v))),
		}
		ins = append(ins, in)
	}
	sort.Slice(ins, func(i, j int) bool {
		if ins[i].sore < ins[j].sore {
			return true
		} else if ins[i].sore > ins[j].sore {
			return false
		}
		return ins[i].rank < ins[j].rank
	})
	return ins[0].rank
}
