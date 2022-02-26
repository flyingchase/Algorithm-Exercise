package intership

import "math"

// 最长回文子串的长度
func countSubstringsPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, dp := 0, make([][]bool, len(s))
	// dp[i][j]表示从i~j 位置内为回文串
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			// 状态转移方程表示
			// 当 dp[i+1][j-1]内部为回文串且两侧边界外的s[i]==s[j]相同则代表此处i,j 亦为回文
			// 考虑 s[i]==s[j]但长度 j-i<3 三个以内肯定为回文
			dp[i][j] = (s[i] == s[j]) && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] {
				if res < (j - i + 1) {
					res = j - i + 1
				}
			}
		}
	}
	return res
}

// 最大子数组之和
func maxSubArraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i]为以 i 为结束位置的最大子数组之和
	length, dp := len(nums), make([]int, len(nums))
	for i := 0; i < length; i++ {
		dp[i] = math.MinInt32
	}
	// 边界 dp[0]为第一个 num
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < length; i++ {
		if nums[i] <= dp[i-1]+nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res <= dp[i] {
			res = dp[i]
		}
	}
	return res
}
