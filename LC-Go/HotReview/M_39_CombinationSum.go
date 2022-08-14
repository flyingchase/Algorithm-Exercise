package HotReview

func combinationSum(candidates []int, target int) [][]int {

	if candidates == nil {
		return [][]int{}
	}
	sum, cur, res, index := 0, []int{}, [][]int{}, 0
	dfsCombinationSum(candidates, sum, cur, &res, index, target)
	return res

}

func dfsCombinationSum(nums []int, sum int, cur []int, res *[][]int, index int, target int) {
	if sum == target {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		sum += nums[i]
		cur = append(cur, nums[i])
		dfsCombinationSum(nums, sum, cur, res, i, target)
		cur = cur[:len(cur)-1]
	}

}
