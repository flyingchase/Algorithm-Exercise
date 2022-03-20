package review

import (
	intership "22intership"
	"math"
	"sort"
)

type TreeNode = intership.TreeNode

// 接雨水 46
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	l, r := 0, len(height)-1
	res := 0
	pl, pr := height[l], height[r]
	for l < r {
		if height[l] < height[r] {
			if height[l] < pl {
				res += pl - height[l]
			} else {
				pl = height[l]
			}
			l++
		} else {
			if height[r] < pr {
				res += pr - height[r]
			} else {
				pr = height[r]
			}
			r--
		}
	}
	return res
}

// 64-最小路径和
// 矩阵之间移动只能往下往右
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 236 lca二叉树两节点的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l != nil {
		if r != nil {
			return root
		}
		return l
	}
	return r
}

// 47-非重复的全排列
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	dfsPermuteUnique(nums, []int{}, 0, &res, &visited)
	return res
}
func dfsPermuteUnique(nums []int, cur []int, index int, res *[][]int, visited *[]bool) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && (*visited)[i-1]) {
			continue
		}
		cur = append(cur, nums[i])
		(*visited)[i] = true
		dfsPermuteUnique(nums, cur, index+1, res, visited)
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}

// 340 至多含有 k个不同字符的最长子串
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) < k {
		return 0
	}
	l, r, res := 0, 0, 0
	dict := map[byte]int{}
	for r < len(s) {
		dict[s[r]]++
		for len(dict) > k {
			dict[s[l]]--
			if dict[s[l]] == 0 {
				delete(dict, s[l])
			}
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}

// 字符串转为整数
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 1
	index := 0
	for ; index < len(s); index++ {
		if s[index] == ' ' {
			continue
		}
	}
	str := []byte(s)
	if str[index] == '-' {
		sign = -1
		index++
	}
	for ; index < len(s); index++ {
		res += res*10 + int(str[index]-'0')
		if res > math.MaxInt32 {
			return -1
		}
	}
	return res * sign
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	return judge(left, right)
}
func judge(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if r.Val != r.Val || r == nil || l == nil {
		return false
	}
	return judge(l.Left, r.Right) && judge(l.Right, r.Left)
}
func characterReplacement(s string, k int) int {

}
