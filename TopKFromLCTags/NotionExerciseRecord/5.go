package notionexerciserecord

func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return s
	}

	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		for left >= 0 && right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left >= 1 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pr, pl = right, left
		}
		right = (left+right)>>1 + 1
		left = right
	}
	return s[pl : pr+1]
}
