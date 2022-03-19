package days

import (
	"sort"
	"strconv"
)

// 5-最长回文子串
// 中心轴对称法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = axial(s, i, i, res)
		res = axial(s, i, i+1, res)
	}
	return res
}
func axial(s string, l, r int, res string) string {
	sub := ""
	for r < len(s) && l >= 0 && s[l] == s[r] {
		sub = s[l : r+1]
		l--
		r++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}

// 最长回文子串
// 滑动窗口
func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	l, r, pl, pr := 0, -1, 0, 0
	for l < len(s) {
		// 移至重复字符边界
		for r < len(s)-1 && s[l] == s[r+1] {
			r++
		}
		// 找回文
		for l >= 1 && r < len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if pr-pl < r-l {
			pr, pl = r, l
		}
		l = (r+l)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}

// dp[i][j]means index from i to j is longestPalindrome
func longestPalindrome3(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			// 特例排除
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || (j-i+1) > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				// dfsNumIslands(grid, i, j, res)
				bfsNumislands(&grid, i, j)
			}
		}
	}
	return res
}
func dfsNumIslands(matrix [][]byte, i, j int, res int) {
	if i >= len(matrix) || j >= len(matrix[0]) || i < 0 || j < 0 || matrix[i][j] == '0' {
		return
	}
	matrix[i][j] = '0'
	dfsNumIslands(matrix, i+1, j, res)
	dfsNumIslands(matrix, i-1, j, res)
	dfsNumIslands(matrix, i, j+1, res)
	dfsNumIslands(matrix, i, j-1, res)
}
func bfsNumislands(matrix *[][]byte, row, col int) {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := make([][]int, 0)
	queue = append(queue, []int{row, col})
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			size--
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]
			if x < 0 || y < 0 || x >= len(*matrix) || y >= len((*matrix)[0]) || (*matrix)[x][y] == '0' {
				continue
			}
			(*matrix)[x][y] = '0'
			for _, dir := range dirs {
				queue = append(queue, []int{x + dir[0], y + dir[1]})
			}
		}
	}
}

// 全排列
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	sort.Ints(nums)
	backtrackPermute(nums, &res, []int{})
	return res
}
func backtrackPermute(nums []int, res *[][]int, cur []int) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrackPermute(nums, res, cur)
		cur = cur[:len(cur)-1]
	}
}

// 字符串相加
func addStrings(num1 string, num2 string) string {
	// 双指针，逆序-进位-反转
	res := ""
	i, j, carry := len(num1)-1, len(num2)-1, 0
	var n1, n2, sum int
	for i >= 0 || j >= 0 {
		n1, n2 = 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		// 该位权数和
		sum = n1 + n2 + carry
		// 进位
		carry = sum / 10
		// 每次的res都是逆序的，所以最后是正序
		res = strconv.Itoa(sum%10) + res
		i--
		j--
	}
	if carry > 0 {
		return "1" + res
	}
	return res
}
