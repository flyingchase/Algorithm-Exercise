package exercise

import "sort"

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res, list, index := [][]int{}, []int{}, 0
	sort.Ints(nums)
	dfs(&res, list, index, nums)
	return res
}

func dfs(res *[][]int, list []int, index int, nums []int) {
	*res = append(*res, append([]int{}, list...))
	for i := index; i < len(nums); i++ {
		if i > index && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		dfs(res, list, i+1, nums)
		list = list[:len(list)-1]
	}
}
