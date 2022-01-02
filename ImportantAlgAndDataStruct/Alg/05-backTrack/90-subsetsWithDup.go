package backtrack

import "sort"

// M-90-子集
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	sort.Ints(nums)
	batcktrackSubsSet(nums, &res, []int{}, 0)
	return res
}
func batcktrackSubsSet(nums []int, res *[][]int, temp []int, index int) {
	*res = append(*res, append([]int{}, temp...))
	for i := index; i < len(nums); i++ {
		if i > index && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		// recutsion
		batcktrackSubsSet(nums, res, temp, i+1)
		// 剪枝
		temp = temp[:len(temp)-1]
	}
}
