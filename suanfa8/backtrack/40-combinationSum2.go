package backtrack

import "sort"

// 子集合
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)
	backtrackSubsetsWithDup(nums, []int{}, &res, 0)
	return res
}
func backtrackSubsetsWithDup(nums, cur []int, res *[][]int, index int) {
	*res = append(*res, append([]int{}, cur...))
	for i := index; i < len(nums); i++ {
		// 去重
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		backtrackSubsetsWithDup(nums, cur, res, i+1)
		cur = cur[:len(cur)-1]
	}
}
