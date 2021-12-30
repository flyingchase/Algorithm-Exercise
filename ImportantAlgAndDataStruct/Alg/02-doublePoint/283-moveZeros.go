package doublepoint

// 将数组中的 0 全部移动到末尾
// 保持非 0 元素的相对位置
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	for count > 0 {
		count--
		nums = append(nums, 0)
	}
}
func moveZeroes_TwoPoints(nums []int) {
	left, N := 0, len(nums)
	// right 将所有非零均前移，覆盖
	for right := 0; right < N; right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	// left 将已经移动的部分全置为 0
	for left < N {
		nums[left] = 0
		left++
	}
}
