package binarysearch

import "sort"

// M-1283-使结果不超过阈值的最小除数
// wrong now waitting to solve
func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, threshold
	curSum := len(nums)
	sort.Ints(nums)
	right = nums[len(nums)-1]
	for left <= right {
		mid := left + (right-left)>>1
		tempSum := validSmallestDivisor(nums, mid)
		if curSum > tempSum {
			right = mid - 1
		} else {
			curSum = tempSum
			left = mid + 1
		}
	}
	return left
}
func validSmallestDivisor(nums []int, divisor int) int {
	sum := 0
	for _, num := range nums {
		temp := num / divisor
		if num%divisor != 0 {
			temp++
		}
		sum += temp
	}
	return sum
}
