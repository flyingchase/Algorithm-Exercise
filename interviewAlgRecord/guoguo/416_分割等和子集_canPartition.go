package guoguo

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}
	n, halfSum := len(nums), sum>>1

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, halfSum+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < halfSum+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else if j == 0 {
				dp[i][j] = true
			} else {
				if j >= nums[i-1] {
					dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[n][halfSum]
}
