package interviewoffer

import "math"

func printNumbers(n int) []int {
	res := []int{}
	max := int(math.Pow(10.0, float64(n)))
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	return res
}
