package selections

// 大于制定数的最短子数组的长度
func minSubArrayLen(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r, sum, res := 0, 0, 0, len(nums)+1
	for r <= len(nums) && l <= r {
		if sum < target {
			// 右边走到长度时候
			if r < len(nums) {
				sum += nums[r]
			}
			r++
		} else {
			// shrink
			if res > r-l {
				sum -= nums[l]
				res = r - l
				l++
			}
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}
