package fifthround

func longestpalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {

		res = maxPalindrome(s, i, i)

	}
	return res
}
