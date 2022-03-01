package intership

// 判断是否为单调队列
// 使用位运算&=
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return false
	}
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		inc = A[i-1] <= A[i]
		dec = A[i-1] >= A[i]
	}
	return inc || dec
}

// 11 数组矩形的最大面积
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		h := Min(height[left], height[right])
		width := right - left
		res = Max(res, h*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 接雨水
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
