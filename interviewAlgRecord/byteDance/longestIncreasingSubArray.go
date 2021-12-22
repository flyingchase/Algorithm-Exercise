package bytedance

// 最长上升子串
// 暴力搜索法
func longestIncreasingSubArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	res := []int{}
	for i := 0; i < len(nums)-1; i++ {
		temp := []int{nums[i]}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[j-1] {
				temp = append(temp, nums[j])
			} else {
				break
			}
		}
		if len(temp) > len(res) {
			res = temp
		}
	}
	return res
}
