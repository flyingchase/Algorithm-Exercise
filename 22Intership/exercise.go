package intership

import (
	"math"
	"sort"
)

//  修改字符出现最多的次数的 num
func maxFrequency1838(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	countMap, res := make(map[int]int, 0), 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		count := countMap[cur]
		if i > 0 {
			operationCount := k
			for j := i - 1; j >= 0; j-- {
				tmp := nums[j]
				diff := cur - tmp
				if operationCount >= diff {
					add := operationCount / diff
					min := math.Min(float64(countMap[tmp]), float64(add))
					operationCount -= int(min) * (diff)
					count += int(min)
				} else {
					break
				}
			}
		}
		if res <= count {
			res = count
		}
	}
	return res
}
