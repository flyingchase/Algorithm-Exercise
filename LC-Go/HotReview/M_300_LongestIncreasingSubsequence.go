package HotReview

import "sort"

/*
LIS 最长上升子序列

*/
func lengthOfLIS(nums []int) int {

	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max300(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max300(res, dp[i])
	}
	return res
}
func max300(j int, i int) int {
	if i > j {
		return i
	} else {
		return j
	}

}
func lengthOfLIS1(nums []int) int {

	dp := []int{}

	for _, num := range nums {

		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {

			dp[i] = num

		}

	}
	return len(dp)

}
