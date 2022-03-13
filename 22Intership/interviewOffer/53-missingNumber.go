package interviewoffer

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := (nums[0] + nums[len(nums)-1]) * (len(nums)) / 2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if res < sum {
		return sum - res
	}
	return res - sum
}
