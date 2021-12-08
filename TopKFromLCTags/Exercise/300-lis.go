package exercise

import "sort"

// 最长上升子序列
// sort.SearchInts 查找 index
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		index := sort.SearchInts(dp, num)
		if len(dp) == index {
			dp = append(dp, num)
		} else {
			dp[index] = num
		}
	}
	return len(dp)
}

// lis
// 逐个比价
func lengthOfLIS2(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
				}
			}
		}
		dp[i] = dp[i] + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
