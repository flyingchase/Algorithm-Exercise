package slidingwindow

import "math"

// 424. 替换后的最长重复字符
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	count, left, res := make([]int, 0, 26), 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		for right-left+1-maxInCount(count) > k {
			count[s[left]-'A']--
			left++
		}
		res = int(math.Max(float64(res), float64(right-left+1)))
	}
	return res
}
func maxInCount(count []int) int {
	if len(count) == 0 {
		return 0
	}
	max := count[0]
	for i := 1; i < len(count); i++ {
		if max < count[i] {
			max = count[i]
		}
	}
	return max
}
