package HotReview

// solution1 N^2 N^2
// 动态规划 dp[i][j]为字符串 i 到 j 个字符的子串是否是回文串
// dp[i][j]=(s[i]--s[j])&&(j-i<3)||dp[i+1][j-1]
// j-i=1;j-i=2为特殊情况
func longestPalindrome1(s string) string {
	res, dp := "", make([][]bool, len(s))
	// initialize
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i < 3) || dp[i+1][j-1]
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				// update res
				// condition: res==""||newRes'len>oldRes'len
				res = s[i : j+1]
			}
		}
	}
	return res
}

/*
solution2 中心扩散法 N^2 N^1
dp中任意起始和中止范围内的字符串判断 不是最长回文字符串则无需判断

- 找到轴心（偶数中心虚拟 ）
- 枚举每个轴线位置
	- 假设为最长回文子串为偶数则虚拟中心扩散 假设奇数则正中心两边扩散
	- 扩散中判断是否对称

*/
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// two assumptions from the center the sides
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i int, j int, res string) string {

	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		// move to the sides require the palindrome
		i--
		j++
	}

	if len(res) < len(sub) {
		return sub
	}
	return res
}

/*
solution3 Sliding Windows


*/
func longestPalindrome3(s string) string {

	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 回文边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		// 暂存目前的边界
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置回文中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
