package guoguo

func backpack(weight, value []int, cap int) int {
	// dp[i][x]=y means 在前 i 件物品时容量为 x 下的最大价值
	// dp[N+1][W+1] N为物品数量, W为最大容量
	// i 物品种类，j 需求
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 1; i <= len(weight); i++ {
		for j := 1; j <= cap; j++ {
			// validate
			if weight[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i-1])
			}
		}
	}
	return dp[len(weight)][cap]
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
