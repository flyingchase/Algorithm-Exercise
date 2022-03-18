package backtrack

import "sort"

// 组合总数
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	res, visited := [][]int{}, make([]bool, len(candidates))
	backtrackCombinationSum(candidates, []int{}, &res, &visited, target, 0)
	return res
}
func backtrackCombinationSum(nums, cur []int, res *[][]int, visited *[]bool, carry, index int) {
	// 优化，carry 时超时
	if carry < 0 {
		return
	}
	if carry == 0 {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := index; i < len(nums); i++ {
		if (*visited)[i] == true || i > index && nums[i] == nums[i-1] && !(*visited)[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		(*visited)[i] = true
		backtrackCombinationSum(nums, cur, res, visited, carry-nums[i], i+1)
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}
