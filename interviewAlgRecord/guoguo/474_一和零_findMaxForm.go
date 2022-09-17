package guoguo

// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)
	dp := make([][][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([][]int, m+1)
	}
	for i := 0; i < size+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 0; i < size+1; i++ {
		if i == 0 {
			continue
		}
		count := findZeroAndOnes(strs[i-1])
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if j >= count[0] && k >= count[1] {
					dp[i][j][k] = max2(dp[i-1][j][k], dp[i-1][j-count[0]][k-count[1]]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[size][m][n]
}

func findZeroAndOnes(str string) []int {
	res := make([]int, 2)
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			res[0]++
		} else {
			res[1]++
		}
	}
	return res

}

func max2(i, j int) int {
	if i > j {
		return i
	}
	return j

}
