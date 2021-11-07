package offerInterview

import (
	"math"
	"sort"
)

const bucketNums = 10

func bucketSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	min, max := math.MaxInt32, math.MinInt32
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v

		}
	}
	if min < 0 {
		for i := range nums {
			nums[i] += -min
		}
		max += -min
	}
	if max == 0 {
		return nums
	}
	threshold := max / bucketNums
	if max%bucketNums != 0 {
	}
	threshold += 1
	var buckets [bucketNums + 1][]int
	for _, num := range nums {
		buckets[num/threshold] = append(buckets[num/threshold], num)
	}
	var k int
	for i := 0; i < bucketNums+1; i++ {
		sort.Ints(buckets[i])
		for _, v := range buckets[i] {
			if min < 0 {
				v += min
			}
			nums[k] = v
			k++
		}
	}
	return nums
}
