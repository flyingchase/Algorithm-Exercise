package firstround

/*
* 区间合并，将输入数组的重叠区间合并并输出
* */

type Interval struct {
	Start int
	End   int
}

func merge(intervals [][]int) [][]int {
	interval := make([]Interval, len(intervals))
	// 将 数组对转化为自定义的数据结构存储区间的边界
	for index, num := range intervals {
		interval[index].Start = num[0]
		interval[index].End = num[1]
	}
	resMerge := mergeIntercals(interval)
	res := [][]int{}
	// 将结构体解耦为数组
	for _, num := range resMerge {
		cur := make([]int, 2)
		cur[0] = num.Start
		cur[1] = num.End
		res = append(res, cur)
	}
	return res
}

// 合并重叠区间
func mergeIntercals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	curIndex := 0
	// res 逐个与已排序完毕的结构体数组比较，起点小于待比较intervals的终点则追加
	// 起点大于待比较intervals的终点则取两者终点的最大值
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > res[curIndex].End {
			curIndex++
			res = append(res, intervals[i])
		} else {
			if intervals[i].End > res[curIndex].End {
				res[curIndex].End = intervals[i].End
			}
		}
	}
	return res
}

// 按照 intervals.Start 排序,Start相同则按End 排序
func quickSort(intervals []Interval, l, r int) {
	if l >= r {
		return
	}
	p := paratition(intervals, l, r)
	quickSort(intervals, l, p[0]-1)
	quickSort(intervals, p[1]+1, r)
}
func paratition(intervals []Interval, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if intervals[l].Start < intervals[r].Start || (intervals[l].Start == intervals[r].Start && intervals[l].End < intervals[r].End) {
			less++
			intervals[l], intervals[less] = intervals[less], intervals[l]
			l++
		} else if intervals[l].Start > intervals[r].Start || (intervals[l].Start == intervals[l].Start && intervals[l].End > intervals[r].End) {
			more--
			intervals[l], intervals[r] = intervals[r], intervals[l]
		} else {
			l++
		}
	}
	intervals[more], intervals[r] = intervals[r], intervals[more]
	return []int{less + 1, more}
}
