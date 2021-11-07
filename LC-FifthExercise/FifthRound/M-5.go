package fifthround

func longestpalindrome(s string) string {
	// dp means [][]from i to j has Palidrome
	if len(s) == 0 {
		return ""
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = ((j-i > 3) || dp[i+1][j-1]) && s[i] == s[j]
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func longestpalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(s); i++ {

		res = maxPalidrome(s, res, i, i)
		res = maxPalidrome(s, res, i, i+1)
	}

	return res
}

func maxPalidrome(s string, res string, i, j int) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
func longestpalindrome3(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right] {
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
