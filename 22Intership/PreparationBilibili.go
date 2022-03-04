package intership

// 01 矩阵
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	m, n := len(mat), len(mat[0])
	queue := make([][]int, 0)
	// 将所有 0 入队且 1 设为-1 标志为未访问过的 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	dirX, dirY := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	// 从 0 开始宽度搜索
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		// 取出队首所记录的下标，开始访问
		for i := 0; i < 4; i++ {
			newX, newY := x+dirX[i], y+dirY[i]
			// 四周为-1 则代表为未访问过的 1
			// 更新该点到 0 的距离为 mat[x][y]+1
			if newX >= 0 && newX < m && newY >= 0 && newY < n && mat[newX][newY] == -1 {
				mat[newX][newY] = mat[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	return mat
}
