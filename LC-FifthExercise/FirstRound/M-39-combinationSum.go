package firstround

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if candidates == nil {
		return res
	}
	sum, cur, index := 0, []int{}, 0
	sort.Ints(candidates)
	dfsCombinationSum(candidates, sum, cur, &res, index, target)
	return res
}
func dfsCombinationSum(candidates []int, sum int, cur []int, res *[][]int, index, target int) {
	if sum >= target {
		if sum == target {
			*res = append(*res, append([]int{}, cur...))
			// tmp := make([]int, len(cur))
			// copy(tmp, cur)
			// *res = append(*res, tmp)
			// return
			// *res=append(*res, append(make([]int, 0,len(cur)), cur...))

		} else {
			return
		}
	}
	for i := index; i < len(candidates); i++ {
		sum += candidates[i]
		cur = append(cur, candidates[i])
		dfsCombinationSum(candidates, sum, cur, res, i, target)
		cur = cur[:len(cur)-1]
	}
}
