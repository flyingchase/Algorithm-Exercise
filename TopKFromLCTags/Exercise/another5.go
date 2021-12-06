package exercise

func longestPalindrome02(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalidrome02(s, i, i, res)
		res = maxPalidrome02(s, i, i+1, res)
	}
	return res
}

func maxPalidrome02(s string, left int, right int, res string) string {
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
func longestPalindrome03(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
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
			pr, pl = right, left
		}
		left = (right+left)>>1 + 1
		right = left
	}
	return res
}
