package thirdround

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i > 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] == (s[i] == s[j]) && ((j-i > 3) || dp[i-1][j+1])
		}

	}

}
