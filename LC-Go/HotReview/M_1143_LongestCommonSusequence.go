package HotReview

/*
两字符串中的最长公共子序列
定义 dp[i][j]为长度为 i 的 text1[0:i-1]和长度为 j 的 text2[0:j-1]的最长公共子序列的长度
	- 边界为 0||0---->dp[i][0]||dp[0][j]均为 0（数组初始化多了 1 默认为0 无须再处理）
	- 状态转移方程：dp[i][j]=dp[i-1][j-1]+1	(text1[i-1]==text2[j-1])
				 (text1[i-1]!=text2[j-1])	max(dp[i-1][j],dp[i][j-1])产生
	- 最终为 dp[len(text1)][len(text2)]
时空复杂度均为O(mn)
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	// 有一个 len 为 0 即退出比较
	if len(text1)*len(text2) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i, v1 := range text1 {
		for j, v2 := range text2 {
			dp[i+1][j+1] = max1143(dp[i+1][j], dp[i][j+1])
			if v1 == v2 {
				dp[i+1][j+1] = max1143(dp[i][j]+1, dp[i+1][j+1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
func max1143(i, j int) int {
	if i < j {
		return j
	}
	return i

}
func longestCommonSubsequence2(text1 string, text2 string) int {
	size1, size2 := len(text1), len(text2)
	if size2*size1 == 0 {
		return 0
	}
	lcs := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		lcs[i] = make([]int, size2+1)
		lcs[i][0] = 0
	}
	for i := 0; i <= size2; i++ {
		lcs[0][i] = 0
	}
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max1143(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	return lcs[size1][size2]
}
