package codetop

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := make(map[byte]int, 0)
	res := 0
	l, r := 0, 0
	for r < len(s) {
		if _, ok := count[s[r]]; !ok {
			count[s[r]]++
			r++
		} else {
			delete(count, s[l])
			l++
		}
		if res < r-l+1 {
			res = r - l + 1

		}
	}
	return res
}
