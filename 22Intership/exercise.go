package intership

import (
	"math"
	"sort"
)

//  修改字符出现最多的次数的 num
func maxFrequency1838(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	countMap, res := make(map[int]int, 0), 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		count := countMap[cur]
		if i > 0 {
			operationCount := k
			for j := i - 1; j >= 0; j-- {
				tmp := nums[j]
				diff := cur - tmp
				if operationCount >= diff {
					add := operationCount / diff
					min := math.Min(float64(countMap[tmp]), float64(add))
					operationCount -= int(min) * (diff)
					count += int(min)
				} else {
					break
				}
			}
		}
		if res <= count {
			res = count
		}
	}
	return res
}

func UpdateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}
	m, n := len(mat), len(mat[0])
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dirX, dirY := []int{0, 0, 1, -1}, []int{1, -1, 0, 0}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node[0], node[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+dirX[i], y+dirY[i]
			if mat[newX][newY] == -1 && newX >= 0 && newX < m && newY >= 0 && newY < n {
				mat[newX][newY] = mat[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	return mat
}

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
				bfsSearchNumsIsLands(grid, i, j)
			}
		}
	}
	return res
}
func bfsSearchNumsIsLands(grid [][]byte, row int, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0] || grid[row][col] == '0' {
		return
	}
	(grid)[row][col] = 0
	bfsSearchNumsIsLands(grid, row+1, col)
	bfsSearchNumsIsLands(grid, row-1, col)
	bfsSearchNumsIsLands(grid, row, col+1)
	bfsSearchNumsIsLands(grid, row, col-1)
}

func lca(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l, r := lca(root.Left, p, q), lca(root.Right, p, q)
	if l == nil && r == nil {
		return nil
	}
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}
