package firstround

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 记录当前位置字符的次数，数字+英文字母+符号空格=256
	var freq [256]int
	res, left, right := 0, 0, -1
	for right+1 < len(s) {
		// extend
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			// shrink
			freq[s[left]-'a']--
			left++
		}
		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}
