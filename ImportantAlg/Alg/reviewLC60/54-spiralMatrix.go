package reviewlc60

func spiralMatirx(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	sum, count := m*n, 0
	top, bottom, left, right := 0, m-1, 0, n-1
	res := []int{}
	for count < sum {
		i, j := top, left
		for count < sum && j <= right {
			res = append(res, matrix[i][j])
			j++
			count++
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
		for count < sum && i < top {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return res
}
