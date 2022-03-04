package interviewprepare

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, cur, res := make([]bool, len(nums)), []int{}, [][]int{}
	// 非重复则需要有序
	sort.Ints(nums)
	index := 0
	dfsPermuteUnique(nums, &res, cur, index, &used)
	return res
}
func dfsPermuteUnique(nums []int, res *[][]int, cur []int, index int, used *[]bool) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		// 不重复序列，i已经使用过或者与前一个数相等且前一个数未使用
		if (*used)[i] || ((i > 0) && (nums[i] == nums[i-1]) && !(*used)[i-1]) {
			continue
		}
		(*used)[i] = true
		cur = append(cur, nums[i])
		dfsPermuteUnique(nums, res, cur, index+1, used)
		// 剪枝
		cur = cur[:len(cur)-1]
		(*used)[i] = false
	}
}
