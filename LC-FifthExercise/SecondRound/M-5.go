package secondround

func longestSunstringPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (j-i > 3 || dp[i+1][j-1]) && (s[i] == s[j])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func longestSunstringPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(s); i++ {

		res = maxPalindromeHelper(s, i, i, res)
		res = maxPalindromeHelper(s, i, i+1, res)
	}
	return res
}
func maxPalindromeHelper(s string, i int, j int, res string) string {
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

func longestSunstringPalindrome3(s string) string {
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
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
