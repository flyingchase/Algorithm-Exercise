package sixthround

func longestUniqueSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	left, right, res := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}

		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}
