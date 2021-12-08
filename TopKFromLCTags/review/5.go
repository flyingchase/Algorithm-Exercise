package review

func longestPalindrome5(s string) string {
	if len(s) <= 0 {
		return s
	}

	var res string
	for i := 0; i < len(s); i++ {
		res = maxPalindrome5(s, i, i, res)
		res = maxPalindrome5(s, i, i+1, res)
	}
	return res
}
func maxPalindrome5(s string, left int, right int, res string) string {
	sub := ""
	for left >= 0 && right < len(s) && s[left] == s[right] {
		sub = s[left : right+1]
		left--
		right++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
