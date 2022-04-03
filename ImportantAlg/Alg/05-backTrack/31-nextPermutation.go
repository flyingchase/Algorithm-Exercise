package backtrack

// M-31-下一个排列
func nextpermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 找到第一个降序 number
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	// 不存在更大的排列
	if i < 0 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}
	// 找到第一个比nums[i]稍小的数
	j := i + 1
	for ; j < len(nums) && nums[j] > nums[i]; j++ {
	}
	j--
	nums[i], nums[j] = nums[j], nums[i]
	i++
	// 将 i 之后部分逆序
	for j = len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
