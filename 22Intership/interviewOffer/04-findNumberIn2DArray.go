package interviewoffer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		v := matrix[row][col]
		if target == v {
			return true
		} else if target > v {
			row++
		} else {
			col--
		}
	}
	return false
}
