package topkTags

// 最长公共子串和最长公共子序列

func longestCommonSubstring(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	m, n := len(a), len(b)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i] == a[j] {
				if i >= 1 && j >= 1 {
					// 为二维表格斜上方值加一
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = 0
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

func longestCommonSubsequence(text1, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	res := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					// 相等则斜上角+1
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i > 0 && j > 0 {
					// 不等则正上方和正左方的最大值+1
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				} else if i == 0 && j > 0 {
					// 边界处理
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
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 最短编辑距离问题
// 将一个字符串 a 编辑为另外一个字符串 b 所需的最小编辑次数
// 编辑动作为：替换 删除 插入

/*
1. 在字符串头部插入示空的字符，将 dp 表格为 (m+1)*(n+1)
2. 第一行，第一列则为单字符串不断插入对应的字符
	dp[0][j]=j dp[i][0]=i
3. dp[i][j]为原字符串以 i-1结尾编辑到原字符串 j-1结尾的最小编辑距离
考虑左上方:
	尾字符相等则拷贝左上方值；否则最多左上方+1(直接替换最后一位)
考虑左侧/上侧：
	left=dp[i][j-1]=1  up=dp[i-1][j]+1
4. 会用到左方 上方 左上方值，因此添加空字符串使得第一行和第一列被初始化
5. res:=dp[m][n]
*/
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	m, n := len(word1), len(word2)
	// 假设 word1和 word2已经插入首字符
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 首行不必再考虑
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left, up := dp[i][j-1]+1, dp[i-1][j]+1
			left_up := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_up = dp[i-1][j-1] + 1
			}
			dp[i][j] = min(min(left, up), left_up)
		}
	}
	return dp[m][n]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
