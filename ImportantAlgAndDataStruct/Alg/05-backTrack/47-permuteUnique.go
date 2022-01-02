package backtrack

import "sort"

// M-47-全排列II
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, visited := make([][]int, 0), make([]bool, len(nums))
	sort.Ints(nums)
	backtrackPermuteUnique(nums, &res, &visited, []int{}, 0)
	return res
}
func backtrackPermuteUnique(nums []int, res *[][]int, visited *[]bool, cur []int, index int) {
	// return
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for i := 0; i < len(nums); i++ {

		// 全排列去重
		// 与前一个相等并且前一个未visited||当前指向已visited
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1]) {
			continue
		}
		(*visited)[i] = true
		cur = append(cur, nums[i])
		backtrackPermuteUnique(nums, res, visited, cur, index+1)
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}

}
