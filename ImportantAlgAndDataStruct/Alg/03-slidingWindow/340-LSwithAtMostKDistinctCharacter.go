package slidingwindow

import "math"

// M-340 s 中最长的最多含有 k 个不同字符的子串的长度
// 滑动窗口板子
func lengthofLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	// map 存储每个 byte 出现的次数 故而 map 的 size 需要满足<=k
	m, left, res := make(map[byte]int, 0), 0, 0
	for right := 0; right < len(s); right++ {
		// 进入窗口
		m[s[right]]++
		// 不符合时持续出窗
		for len(m) > k {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}
		// 窗口 valid 时计算
		res = int(math.Max(float64(res), float64(right-left+1)))
	}
	return res
}
