package datawhale50

import (
	"math"
)

func myAtoi(s string) int {
	sign, res, i := 1, 0, 0
	// 删除前导空格
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	// 判断正负
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		res = res*10 + int(s[i]-'0')
		// 数字越界则返回 math.Min/MaxInt32
		if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return sign * res
}
