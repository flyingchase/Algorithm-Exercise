package review

import (
	"math"
)

// 两数组的最长公共子串的长度
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]代表nums1[0:i+1] nums2[0:j+1]之间的最长公共子串的长度
	m, n := len(nums1), len(nums2)
	if m*n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

//76-最小覆盖子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	dict := map[byte]int{}
	var l, r, start, length int
	length = math.MaxInt32
	match := 0
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}
	for r < len(s) {
		if _, ok := dict[s[r]]; ok {
			dict[s[r]]--
			if dict[s[r]] == 0 {
				match++
			}
		}
		r++
		for match == len(dict) {
			if length > r-l {
				length = r - l
				start = l
			}
			if count, ok := dict[s[l]]; ok {
				if count == 0 {
					match--
				}
				dict[s[l]]++
			}
			l++
		}
	}
	if length == math.MaxInt32 {
		return ""

	}
	return s[start : start+length]
}

// 113 二叉树中路径总和等于给定targetSum的所有路径
// 必须根到叶子结点
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	dfsPathSum(&res, root, targetSum, 0, []int{})
	return res
}

func dfsPathSum(res *[][]int, node *TreeNode, target, curSum int, curPath []int) {
	curPath = append(curPath, node.Val)
	curSum += node.Val
	if curSum == target && node.Left == nil && node.Right == nil {
		*res = append(*res, append([]int{}, curPath...))
		return
	}
	if node.Left != nil {
		dfsPathSum(res, node.Left, target, curSum, curPath)
	}
	if node.Right != nil {
		dfsPathSum(res, node.Right, target, curSum, curPath)
	}
	curPath = curPath[:len(curPath)-1]
}

// 437 路径总和，任意结点之间，必须向下即可
// 前缀和问题
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		return 1
	}
	res := 0
	prefixSum := map[int]int{}
	prefixSum[0] = 1
	dfsPathSum2(&res, root, targetSum, 0, []int{}, prefixSum)
	return res
}

func dfsPathSum2(res *int, node *TreeNode, target, curSum int, curPath []int, prefixSum map[int]int) {
	if node == nil {
		return
	}
	curSum += node.Val
	curPath = append(curPath, node.Val)

	*res += prefixSum[curSum-target]
	prefixSum[curSum]++
	dfsPathSum2(res, node.Left, target, curSum, curPath, prefixSum)
	dfsPathSum2(res, node.Right, target, curSum, curPath, prefixSum)
	prefixSum[curSum]--
}
