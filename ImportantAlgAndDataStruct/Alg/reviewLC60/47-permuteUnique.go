package reviewlc60

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	backtrackPermuteUnique(nums, &res, []int{}, 0, &visited)
	return res
}
func backtrackPermuteUnique(nums []int, res *[][]int, cur []int, index int, visited *[]bool) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i]) {
			continue
		}
		(*visited)[i] = true
		cur = append(cur, nums[i])
		backtrackPermuteUnique(nums, res, cur, index+1, visited)
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}
