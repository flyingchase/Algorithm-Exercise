package days

import "math"

// 查找出现的首个位置和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			i, j := mid, mid
			for i >= -1 && j <= len(nums) && (i >= 0 && nums[i] == target || j < len(nums) && nums[j] == target) {
				if i >= 0 && nums[i] == target {
					i--
				}
				if j < len(nums) && nums[j] == target {
					j++
				}
				if i >= 0 && j < len(nums) && nums[i] != target && nums[j] != target {
					break
				}
			}
			return []int{i + 1, j - 1}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}

// 不同的路径
func UniquePath(m, n int) int {
	// dp数组初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 遍历顺序
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 递推公式
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 寻找峰值
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	tmp := math.MinInt32
	nums = append([]int{tmp}, nums...)
	nums = append(nums, tmp)
	l, r := 1, len(nums)-2
	lDict := map[int]struct{}{}
	rDict := map[int]struct{}{}
	for l <= r {
		if nums[l-1] < nums[r] {
			lDict[l] = struct{}{}
		}
		l++
	}
	for r >= 1 {
		if nums[r+1] < nums[r] {
			rDict[r] = struct{}{}
			if _, exist := lDict[r]; exist {
				return r - 1
			}
		}
		r--
	}
	return -1
}

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

// lcp 求两个字符串str1和str2的最长公共前缀
func lcp(str1, str2 string) string {
	length := Min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
