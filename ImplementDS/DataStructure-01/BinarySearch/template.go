package binarysearch

func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
