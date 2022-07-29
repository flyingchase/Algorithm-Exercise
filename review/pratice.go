package review

func twoSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		sum := nums[i] + nums[i+1]
		if sum == target {
			res = append(res, []int{nums[i], nums[i+1]})
		} else if sum > target {

		} else {

		}
	}
	return res
}
