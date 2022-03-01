package intership

import "sort"

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

// M-90-子集
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	// 非重复前提为有序
	sort.Ints(nums)
	batcktrackSubsSet(nums, &res, []int{}, 0)
	return res
}
func batcktrackSubsSet(nums []int, res *[][]int, temp []int, index int) {
	*res = append(*res, append([]int{}, temp...))
	// 子集需要从 index 开始，backtrack 需要从i+1
	for i := index; i < len(nums); i++ {
		if i > index && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		// recutsion 从 i+1 而非 index+1 开始
		batcktrackSubsSet(nums, res, temp, i+1)
		// 剪枝
		temp = temp[:len(temp)-1]
	}
}

// 非重复的全排列
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, visited := make([][]int, 0), make([]bool, len(nums))
	// 非重复前提为有序
	sort.Ints(nums)
	backtrackPermuteUnique(nums, &res, &visited, []int{}, 0)
	return res
}
func backtrackPermuteUnique(nums []int, res *[][]int, visited *[]bool, cur []int, index int) {
	// return
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// 全排列去重
		// 与前一个相等并且前一个未visited||当前指向已visited
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1]) {
			continue
		}
		(*visited)[i] = true
		cur = append(cur, nums[i])
		// backtrack 从 index+1 开始
		backtrackPermuteUnique(nums, res, visited, cur, index+1)
		// 剪枝 注意(*visited)[i]==false
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}

// 下一个稍大的排列
// M-31-下一个排列
func nextpermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 找到第一个降序 number
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	// 不存在更大的排列
	if i < 0 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}
	// 找到第一个比nums[i]稍小的数
	j := i + 1
	for ; j < len(nums) && nums[j] > nums[i]; j++ {
	}
	j--
	nums[i], nums[j] = nums[j], nums[i]
	i++
	// 将 i 之后部分逆序
	for j = len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
