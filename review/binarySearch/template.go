package binarysearch

func template(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := end + (start-end)>>1
		if nums[mid] == target {
			panic("find index at mid is target")
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
