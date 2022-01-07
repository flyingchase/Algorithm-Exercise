package reviewlc60

func rotate(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}
	return matrix
}
