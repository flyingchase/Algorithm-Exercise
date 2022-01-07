package reviewlc60

import "sort"

func subSet(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}
	res := [][]int{}
	sort.Ints(nums)
	dfsSubSet(nums, &res, []int{}, 0)
	return res
}
func dfsSubSet(nums []int, res *[][]int, cur []int, index int) {
	*res = append(*res, append([]int{}, cur...))
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		dfsSubSet(nums, res, cur, i+1)
		cur = cur[:len(cur)-1]
	}
}
