package interviewprepare

// 接雨水
// 双指针，前后往中间靠，不断更新 left 和 right 并求值存入 res
func trap(height []int) int {
	var left, right, res int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			if height[i] < left {
				res += left - height[i]
			} else {
				left = height[i]
			}
			i++
		} else {
			if height[j] < right {
				res += right - height[j]
			} else {
				right = height[j]
			}
			j--
		}
	}
	return res
}
