package topselection

func search(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}
	return searchMid(nums, 0, len(nums)-1, target)
}
func searchMid(nums []int, l, r int, target int) bool {

	for l <= r {

		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		}
		if nums[l] < nums[mid] {
			if target < nums[mid] && target >= nums[l] {
				r = mid
			} else {
				l = mid + 1
			}
		} else if nums[l] > nums[mid] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			l++
		}
	}
	return false
}
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	cout := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cout++
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	for i := 0; i < cout; i++ {
		nums = append(nums, 0)
	}
}
