package exercise

// 最长回文子串中心扩散法
func longestPalindrome(s string) string {

	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var res string
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}
func maxPalindrome(s string, left int, right int, res string) string {

	sub := ""
	for left >= 0 && right < len(s) && s[left] == s[right] {
		sub = s[left : right+1]
		left--
		right++
	}
	if len(res) < len(sub) {
		res = sub
	}
	return res
}

// 最长回文子串———滑动窗口法
func longestPalindrome2(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)>>1 + 1
		right = left
	}
	return s[pl : pr+1]
}
