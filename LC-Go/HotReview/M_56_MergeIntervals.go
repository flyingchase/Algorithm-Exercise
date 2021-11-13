package HotReview

/*
区间合并
将重叠的区间合并为一个区间
*/
type Interval struct {
	Start int
	End   int
}

func merge(intervals [][]int) [][]int {
	interval := make([]Interval, len(intervals))
	for index, num := range intervals {
		interval[index].Start = num[0]
		interval[index].End = num[1]
	}

	resMerge := mergeIntervals(interval)
	res := make([][]int, 0)
	for _, num := range resMerge {
		cur := make([]int, 2)
		cur[0] = num.Start
		cur[1] = num.End
		res = append(res, cur)
	}
	return res
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	// 手写快排 使得依照区间的起点排序 起点相同按照终点排序
	quicksort(intervals, 0, len(intervals)-1)
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > res[curIndex].End {
			curIndex++
			res = append(res, intervals[i])
		} else {
			res[curIndex].End = max56(intervals[i].End, res[curIndex].End)
		}

	}
	return res
}

func max56(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func quicksort(a []Interval, l int, r int) {
	if l >= r {
		return
	}

	p := paratitions(a, l, r)
	quicksort(a, l, p-1)
	quicksort(a, p+1, r)
}

func paratitions(a []Interval, l int, r int) int {
	pivot := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
