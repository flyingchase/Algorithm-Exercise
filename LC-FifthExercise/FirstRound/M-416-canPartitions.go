package firstround

import (
	"sort"
)

func canParatition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sort.Ints(nums)
	var DFS416 func(int, int, int) bool
	DFS416 = func(i int, sum1, sum2 int) bool {
		if sum1 == sum/2 {
			return true
		}
		if i < 0 || sum1 > sum/2 || sum2 > sum/2 {
			return false
		}
		return DFS416(i-1, sum1+nums[i], sum2) || DFS416(i-1, sum1, sum2+nums[i])
	}
	return DFS416(len(nums)-1, 0, 0)
}
