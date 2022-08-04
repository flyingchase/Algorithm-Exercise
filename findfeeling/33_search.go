package findfeeling

// binarySearch means nlogn can return the index of target in nums
// require nums is ordered
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		// -1 means not found
		return -1
	}
	l, r := 0, len(nums)
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
