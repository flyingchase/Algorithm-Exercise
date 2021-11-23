package topselection

func sortColors(nums []int) {
	m := make(map[int]int, 3)
	for _, num := range nums {
		m[num]++
	}
}
func deleteDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			// 连续多个重复时不跳过(>=3)
			i--
		}
	}
	return len(nums)
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 2 {
		return len(nums)
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
