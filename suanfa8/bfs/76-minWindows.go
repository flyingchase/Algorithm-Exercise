package bfs

import "math"

// 最小子串覆盖
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	l, r, start, length, match := 0, 0, 0, 0, 0
	length = math.MaxInt32
	dict := map[byte]int{}
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
