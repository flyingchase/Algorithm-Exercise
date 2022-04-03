package binarysearch

// 二分查找板子
// 非递归
func binarysearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
	// (end,start)
	// 相等情况直接返回 mid
}

// 二分查找板子
// 递归
func binarysearch2(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)>>1
	if nums[mid] > target {
		return binarysearch2(nums, left, mid-1, target)
	}
	if nums[mid] < target {
		return binarysearch2(nums, mid+1, right, target)
	}
	return mid
}
