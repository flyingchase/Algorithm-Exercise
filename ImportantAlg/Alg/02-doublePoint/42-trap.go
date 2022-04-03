package doublepoint

// 42-接雨水
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	// start end 为左右指针下标遍历
	// l,r 存储高度
	start, end, l, r := 0, len(height)-1, 0, 0
	res := 0
	for start < end {
		if height[start] < height[end] {
			if height[start] < l {
				res += l - height[start]
			} else {
				l = height[start]
			}
			start++
		} else {
			if height[end] < r {
				res += r - height[end]
			} else {
				r = height[end]
			}
			end--
		}
	}
	return res
}
