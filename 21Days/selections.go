package days

import (
	"math"
	"reflect"
)

// M-64-二维网格找左上角到右下角的最小路径
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	// dp[i][j]represent minPathSum index i,j
	// dp为 grid 的副本，累加中+=要求
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = append([]int{}, grid[i]...)
	}
	// board situation
	// 首行为 grid网格横移
	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0]
	}
	// 首列为 grid 网格竖移
	for j := 1; j < n; j++ {
		dp[0][j] += dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 为左侧和上方最小值累加
			dp[i][j] += min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// M-340-至多包含 k 个不同字符的最长子串的长度
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	k %= len(s)
	if len(s) == 1 && k == 1 {
		return 1
	}
	dict, l, r, res := map[byte]int{}, 0, 0, math.MinInt32
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

// M-567 s2中是否包含有 s1 的子串
func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	map1, map2 := make([]int, 26), make([]int, 26)
	// 首个窗口入 map2，同时记录 s1中的字符频次
	for i := 0; i < len1; i++ {
		map1[s1[i]-'a']++
		map2[s2[i]-'a']++
	}
	// 在字符串s2 上从 len1到 len2 滑动窗口
	// 窗口大小固定为 len1
	for i := len1; i < len2; i++ {
		if reflect.DeepEqual(map2, map1) {
			return true
		}
		// 加入新的字符
		map2[s2[i]-'a']++
		// 移出 i-len1 窗口最左侧的字符频次
		map2[s2[i-len1]-'a']--
	}
	return reflect.DeepEqual(map1, map2)
}
