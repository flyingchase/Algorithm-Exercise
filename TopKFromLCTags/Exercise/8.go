package exercise

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var i int
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	var sign int
	if s[i] == '-' {
		sign = -1
	} else if s[i] == '+' {
		sign = 1
	}

	i++
	var res int

	for ; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
		if sign*res < math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * sign

}
