package doublepoint

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		h := minCompare(height[left], height[right])
		width := right - left
		res = maxCompare(res, h*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func minCompare(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func maxCompare(i, j int) int {
	if i > j {
		return i
	}
	return j
}
