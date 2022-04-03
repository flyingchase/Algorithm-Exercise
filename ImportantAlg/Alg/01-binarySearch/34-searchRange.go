package binarysearch

// M-32-在排序数组中查找指定 target 出现的首次和最后一次位置
func searchRange(nums []int, target int) []int {
	res := make([]int, 0, 2)
	res = append(res, -1, -1)
	if len(nums) <= 1 {
		return res
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			start, end := mid, mid
			// 最坏情况下整个数组均为 target
			// 优化可为在 mid 两侧再次二分找到第一个比 mid 小/大的数
			for start >= 1 && nums[start-1] == nums[start] {
				start--
			}
			for end < len(nums)-1 && nums[end] == nums[end+1] {
				end++
			}
			res[0] = start
			res[1] = end
			return res
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

// 三次二分
func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	pivot := -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			pivot = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if pivot == -1 {
		return []int{-1, -1}
	}

	leftBound, RightBound := pivot, pivot
	left, right = 0, pivot
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			leftBound = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left, right = pivot, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			RightBound = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{leftBound, RightBound}

}
