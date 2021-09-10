package Waitting

import "sort"

func subsets(nums []int) [][]int {
	if nums == nil {
		return [][]int{{}}
	}
	cur, res, index := make([]int, 0), make([][]int, 0), 0
	dfsSubsets(nums, cur, &res, index)

	return res
}

func dfsSubsets(nums []int, cur []int, res *[][]int, index int) {

	temp := make([]int, len(cur))
	copy(temp, cur)

	*res = append(*res, temp)

	for i := index; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfsSubsets(nums, cur, res, i+1)
		cur = cur[:len(cur)-1]
	}
}
// 模拟 dfs
func subset1(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, org := range res {
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}



