package datawhale50

func removeDuplicates(nums []int) int {
	for index := 1; index < len(nums); index++ {
		if nums[index] == nums[index-1] {
			// 将 idnex 位置覆盖
			nums = append(nums[:index], nums[index+1:]...)
			index--
		}
	}
	return len(nums)
}
