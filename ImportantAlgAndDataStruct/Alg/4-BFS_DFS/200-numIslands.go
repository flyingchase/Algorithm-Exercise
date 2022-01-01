package bfsdfs

// M-200 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				bfsNumsIslands(&grid, i, j)
				// dfsNumIslands(grid, i, j)
			}
		}
	}
	return res
}

// bfs 宽度优先遍历
func bfsNumsIslands(grid *[][]byte, row, col int) {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := make([][]int, 0)
	queue = append(queue, []int{row, col})
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			size--
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]
			if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) || (*grid)[x][y] == '0' {
				continue
			}
			(*grid)[x][y] = '0'
			for _, dir := range dirs {
				queue = append(queue, []int{x + dir[0], y + dir[1]})
			}
		}
	}
}

// dfs 深度优先遍历
func dfsNumIslands(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dfsNumIslands(grid, row-1, col)
	dfsNumIslands(grid, row+1, col)
	dfsNumIslands(grid, row, col+1)
	dfsNumIslands(grid, row, col-1)
}
