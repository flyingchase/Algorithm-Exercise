package HotReview

import "sort"

/*
无重复数字的全排列
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	used, cur, res := make([]bool, len(nums)), []int{}, [][]int{}
	dfs(nums, 0, cur, &res, &used)
	return res
}

func dfs(nums []int, index int, cur []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		// 新建 tmp 防止指针改变回溯的操作
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			cur = append(cur, nums[i])
			dfs(nums, index+1, cur, res, used)
			// 退回操作
			cur = cur[:len(cur)-1]
			(*used)[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	// 非重复排列前提是 数组有序
	sort.Ints(nums)
	cur, index, res, used := []int{}, 0, [][]int{}, make([]bool, len(nums))
	dfsUnique(nums, &res, cur, index, &used)
	return res
}

func dfsUnique(nums []int, res *[][]int, cur []int, index int, used *[]bool) {
	if index == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 去重关键  i 下边的数已经用过 或者 与前一个用过的数相等且前一个数还没被用过
		if (*used)[i] || ((i > 0) && (nums[i] == nums[i-1]) && !(*used)[i-1]) {
			continue
		}
		(*used)[i] = true
		cur = append(cur, nums[i])
		dfsUnique(nums, res, cur, index+1, used)
		// 去重时候注意 Used变化
		// 剪枝
		cur = cur[:len(cur)-1]
		(*used)[i] = false
	}
}
