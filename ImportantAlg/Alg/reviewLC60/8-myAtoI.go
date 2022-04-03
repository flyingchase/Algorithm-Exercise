package reviewlc60

import "math"

func myAtoI(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return int(s[0] - '0')
	}
	i := 0
	res := 0
	for ; s[i] == ' '; i++ {
	}
	sign := 1
	if s[i] == '-' {
		sign = -1
	} else {
		sign = 1
	}
	i++
	for ; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
		if res > math.MaxInt32 {
			return -1
		}
	}
	return res * sign
}
