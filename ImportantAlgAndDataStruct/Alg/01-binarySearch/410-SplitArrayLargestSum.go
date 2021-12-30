package binarysearch

import (
	"sort"
)

// H-410-分割数组的最大值
// 将整数数组 分割为 m 份，使得 m 份各个子数组之和中的最大值最小
// 二分猜答案
func splitArray(nums []int, m int) int {
	temp := make([]int, 0)
	temp = append(temp, nums...)
	sort.Ints(temp)
	// max 为可分割子数组各自和中的最小值
	// sum 为可分割子数组各自和中的最大值
	// 最为二分的上下界
	max, sum := temp[len(temp)-1], 0
	for _, num := range temp {
		sum += num
	}
	return binary(nums, m, sum, max)
}

// 二分猜答案
func binary(nums []int, m, high, low int) int {
	mid := 0
	for low <= high {
		mid = low + (high-low)>>1
		// 符合分割情况则往中间靠
		// 可将 nums 分割为 m 份，其中有和为 mid
		if valid(nums, m, mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// 判断猜的答案是否符合子数组各自和最大值最小
// 可分割 m 份且每份和不超过subArraySum
func valid(nums []int, m, subArraySum int) bool {
	// count代表分割的分数，一旦 curSum 超过限定的subArraySum，则 count++，同时
	// curSum缩回越界的 num 开始累加和
	curSum, count := 0, 1
	for _, num := range nums {
		curSum += num
		if curSum > subArraySum {
			curSum = num
			count++
			if count > m {
				return false
			}
		}
	}
	return true
}
