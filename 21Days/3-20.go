package days

// spiralOrder
// 模拟螺旋的步进操作即可
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	sum := m * n
	count := 0
	top, bottom, left, right := 0, m-1, 0, n-1
	res := []int{}
	for count < sum {
		i, j := top, left
		for count < sum && j <= right {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top+1, right
		for count < sum && i <= bottom {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i, j = bottom, right-1
		for count < sum && j >= left {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i, j = bottom-1, left
		for count < sum && i > top {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return res
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	return res
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	pl, pr := 0, 0
	for l < r {
	}
}
