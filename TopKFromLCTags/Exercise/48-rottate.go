package exercise

// 旋转图像
// 先对角线转化再垂直轴对称交换
func rotate(matrix [][]int) {
	length := len(matrix)
	// 对角线转换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 垂直轴对称翻转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}
}
