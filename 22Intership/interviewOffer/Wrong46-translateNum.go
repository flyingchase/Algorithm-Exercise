package interviewoffer

import (
	"fmt"
	"strconv"
)

func TranslateNum(num int) int {
	if num >= 0 && num < 10 {
		return 1
	}
	if num >= 10 && num < 26 {
		return 2
	}
	nums := []byte(fmt.Sprint(num))
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(nums); i++ {
		tmp, _ := strconv.Atoi(string(nums[i-2 : i]))
		if tmp >= 10 && tmp < 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)]
}
