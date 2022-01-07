package reviewlc60

import "math"

func minLengthOfSubArraySumGreaterTarget(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	res := math.MaxInt32
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if res > right-left+1 {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	return res
}
