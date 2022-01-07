package reviewlc60

import "alg"

func lis(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for index, _ := range dp {
		dp[index] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = alg.Max(dp[j]+1, dp[i])
			}
			res = alg.Max(res, dp[i])
		}
	}
	return res

}
