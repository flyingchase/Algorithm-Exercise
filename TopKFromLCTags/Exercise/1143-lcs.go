package exercise

// lcs 最长公共子串
func longestCommonSubsequence(text1 string, text2 string) int {
	size1, size2 := len(text1), len(text2)
	if size1*size2 == 0 {

		return 0
	}
	lcs := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		lcs[i] = make([]int, size2+1)
		lcs[i][0] = 0
	}
	for i := 0; i < size2+1; i++ {
		lcs[0][i] = 0
	}
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				if lcs[i-1][j] > lcs[i][j-1] {
					lcs[i][j] = lcs[i-1][j]
				} else {
					lcs[i][j] = lcs[i][j-1]
				}
			}
		}
	}
	return lcs[size1][size2]
}
