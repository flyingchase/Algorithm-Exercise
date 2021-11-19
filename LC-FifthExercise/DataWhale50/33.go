package datawhale50

// 旋转后的排序数组中查找 target
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 二分查找，找到转折的端点
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			// 左端点小于中点
			// 讨论 target 落在 left 和 mid 之间
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右端点小于中点
			// 讨论 target 落在 mid和 right之间
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	// 判断二分后的左端点满足 target 否
	if nums[left] == target {
		return left
	}
	return -1
}
