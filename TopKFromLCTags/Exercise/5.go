package exercise

// 最长回文子串
// 中心轴法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var res string

	for i := 0; i < len(s); i++ {
		res = maxPalidrome(s, i, i, res)
		res = maxPalidrome(s, i, i+1, res)
	}
	return res
}
func maxPalidrome(s string, left, right int, res string) string {
	sub := ""
	for left >= 0 && right < len(s) && s[left] == s[right] {
		sub = s[left : right+1]
		left--
		right++
	}
	if len(sub) >= len(res) {
		return sub
	}
	return res
}

// 最长回文子串
// 滑动窗口
func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	left, right, pr, pl := 0, -1, -1, 0

	for left < len(s) {
		for right+1 < len(s) && s[right+1] == s[left] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			right++
			left--
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (right+left)>>1 + 1
		right = left
	}
	return s[pl : pr+1]
}
