package guoguo

import "math"

func splitArray(nums []int, m int) int {
	copyTmp := make([]int, len(nums))
	copy(copyTmp, nums)
	max := math.MinInt32
	var sum int
	for _, num := range nums {
		sum += num
		if max < num {
			max = num
		}
	}
	return binarySearch(nums, max, sum, m)
}

func binarySearch(nums []int, l, r, m int) int {
	for l <= r {
		mid := l + (r-l)>>1
		if valid(nums, m, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func valid(nums []int, m, subSum int) bool {
	cur, count := 0, 1
	for _, num := range nums {
		cur += num
		if cur > subSum {
			cur = num
			count++
			if count > m {
				return false
			}
		}
	}
	return true
}
