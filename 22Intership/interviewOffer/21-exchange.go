package interviewoffer

func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i]%2 == 0 {
			node := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, node)
			i--
			length--
		}
	}
	return nums
}
