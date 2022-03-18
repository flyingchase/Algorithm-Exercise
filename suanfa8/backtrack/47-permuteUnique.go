package backtrack

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	backtrackPermuteUnique(nums, []int{}, &res, 0, &visited)
	return res
}
func backtrackPermuteUnique(nums, cur []int, res *[][]int, index int, visited *[]bool) {
	if len(nums) == len(cur) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1]) {
			continue
		}
		cur = append(cur, nums[i])
		(*visited)[i] = true
		backtrackPermuteUnique(nums, cur, res, index+1, visited)
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}
