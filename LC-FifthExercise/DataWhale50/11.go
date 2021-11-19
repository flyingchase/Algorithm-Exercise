package datawhale50

import "math"

// 盛水最多的容器
func maxArea(height []int) int {
	// 输入容器只有一边或为空的情况
	if len(height) <= 1 {
		return 0
	}
	var max float64
	i, j := 0, len(height)-1
	// 双指针一边遍历一边更新 area
	for i <= j {
		// 左端值矮
		if height[i] < height[j] {
			max = math.Max(max, float64(height[i]*(j-i)))
			i++
		} else {
			// 右端值矮
			max = math.Max(max, float64(height[j]*(j-i)))
			j--
		}
	}
	return int(max)
}
