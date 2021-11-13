package secondround

type Interval struct {
	Start int
	End   int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	res := [][]int{}
	nums := make([]Interval, len(intervals))
	i := 0
	for _, interval := range intervals {
		nums[i].Start = interval[0]
		nums[i].End = interval[1]
		i++
	}
	resMerge := mergeInterval(nums, 0, len(nums)-1)
	for _, num := range resMerge {
		cur := make([]int, 2)
		cur[0] = num.Start
		cur[1] = num.End
		res = append(res, cur)
	}
	return res
}

func mergeInterval(nums []Interval, l, r int) []Interval {
	if l >= r {
		return nums
	}
	quicksort(nums, l, r)
	res := []Interval{}
	res = append(res, nums[0])
	curIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i].Start > res[curIndex].End {
			curIndex++
			res = append(res, nums[i])
		} else {
			if nums[i].End > res[curIndex].End {
				res[curIndex].End = nums[i].End
			}
		}

	}

	return res
}

func quicksort(nums []Interval, l, r int) {
	if l >= r {
		return
	}
	p := paratitions56(nums, l, r)
	quicksort(nums, l, p[0]-1)
	quicksort(nums, p[1]+1, r)
}
func paratitions56(nums []Interval, l, r int) []int {
	if l >= r {
		return nil
	}
	less, more := l-1, r
	for l < more {
		if nums[l].Start < nums[r].Start || (nums[l].Start == nums[r].Start && nums[l].End < nums[r].End) {
			less++
			nums[less], nums[l] = nums[l], nums[less]
			l++
		} else if nums[l].Start > nums[r].Start || (nums[l].Start == nums[r].Start && nums[l].End > nums[r].End) {
			more--
			nums[more], nums[l] = nums[l], nums[more]
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}
