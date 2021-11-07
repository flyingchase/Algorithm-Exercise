package firstround

func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i < 3) || dp[i+1][j-1]
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}

	}
	return res
}

// 中心扩散法
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 中心轴线假设
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}
func maxPalindrome(s string, i int, j int, res string) string {
	sub := ""
	// 由中心轴线两端移动
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	// 更新 res
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// sliding window
func longestPalindrome3(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		//移动到相同字母最右边
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 回文边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		// 暂存边界
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置回文中心
		left = (left+right)>>1 + 1
		right = left
	}
	return s[pl : pr+1]
}
