package interviewoffer

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res, l, r := 0, 0, 0
	dict := map[byte]int{}
	for r < len(s) {
		if _, ok := dict[s[r]]; !ok {
			dict[s[r]]++
			r++
		} else {
			delete(dict, s[l])
			l++
		}
		if res < r-l {
			res = r - l
		}
	}
	return res
}
