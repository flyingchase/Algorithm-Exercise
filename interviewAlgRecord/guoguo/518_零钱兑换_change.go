package guoguo

// 完全背包 dp[i][j-wt[i-1]]
// 01背包 dp[i-1][j-wt[i-1]]
func change(amount int, coins []int) int {

	size := len(coins)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < size+1; i++ {
		for j := 0; j < amount+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				if j >= coins[i-1] {
					dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[size][amount]
}
