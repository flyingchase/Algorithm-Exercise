package interviewoffer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	count, sum := 0, m*n
	res := []int{}
	top, bottom, left, right := 0, m-1, 0, n-1
	for count < sum {
		i, j := top, left
		for count < sum && j <= right {
			count++
			res = append(res, matrix[i][j])
			j++
		}
		i, j = top+1, right
		for count < sum && i <= bottom {
			count++
			res = append(res, matrix[i][j])
			i++
		}
		i, j = bottom, right-1
		for count < sum && j >= left {
			count++
			res = append(res, matrix[i][j])
			j--
		}
		i, j = bottom-1, left
		for count < sum && i > top {
			count++
			res = append(res, matrix[i][j])
			i--
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return res
}
