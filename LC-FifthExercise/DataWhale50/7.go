package datawhale50

func resverse(x int) int {
	const max = 2147483647
	const min = -2147483648
	var res int
	for x != 0 {
		res = res*10 + x%10
		if res > max || res < min {
			return 0
		}
		x /= 10
	}
	return res
}
