package days

// 括号生成
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	res := []string{}
	dfsGenerateParenthesis(n, n, &res, "")
	return res
}
func dfsGenerateParenthesis(l, r int, res *[]string, cur string) {
	if l == 0 && r == 0 {
		*res = append(*res, cur)
		return
	}
	// 生成左括号
	if l > 0 {
		dfsGenerateParenthesis(l-1, r, res, cur+"(")
	}
	// 括号成对存在，有左括号才会有右括号
	if l < r && r > 0 {
		dfsGenerateParenthesis(l, r-1, res, cur+")")
	}
}

// 最小编辑距离
func MinDistance(word1 string, word2 string) int {
	// dp[i][j]代表以 i-1,j-1 为结尾的word1子串a编辑为word2子串 b 的最小编辑次数
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 边界：将单字符编辑为任意一个i/j 结尾的子串
	// 不断插入即可，则 dp[0][j]=j,dp[i][0]=i
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	// 三种编辑：左侧表格+1，上方表格+1，拷贝左上方或+1即可
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left := dp[i][j-1] + 1
			top := dp[i-1][j] + 1
			left_top := dp[i-1][j-1] + 1
			if word1[i] == word2[j] {
				left_top = dp[i-1][j-1]
			}
			dp[i][j] = minDistanceMin(minDistanceMin(left, top), left_top)
		}
	}
	return dp[m][n]
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m*n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i > 0 && j > 0 {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				} else if i == 0 && j > 0 {
					dp[i][j] = dp[i][j-1]
				} else if i > 0 && j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = 0
				}
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}
