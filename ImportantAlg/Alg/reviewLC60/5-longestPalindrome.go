package reviewlc60

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	pl, pr, l, r := 0, -1, 0, 0
	for l < len(s) {
		for r+1 < len(s) && s[l] == s[r+1] {
			r++
		}
		for l > 0 && r < len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if pr-pl < r-l {
			pl, pr = l, r
		}
		l = (r+l)>>1 + 1
		r = l
	}
	return s[l : r+1]
}
