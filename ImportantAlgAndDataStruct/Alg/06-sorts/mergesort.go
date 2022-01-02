package sorts

func mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mergesortHelper(nums, 0, len(nums)-1)
	return nums
}
func mergesortHelper(nums []int, l, r int) {

	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergesortHelper(nums, l, mid)
	mergesortHelper(nums, mid+1, r)
	merge(nums, l, mid, r)
}
func merge(nums []int, l, mid, r int) {
	pl, pr := l, mid+1
	helper, i := make([]int, r-l+1), 0
	for pl <= mid && pr <= r {
		if nums[pl] <= nums[pr] {
			helper[i] = nums[pl]
			pl++
		} else {
			helper[i] = nums[pr]
			pr++
		}
		i++
	}
	copy(helper[i:],nums[pl:mid+1])
	copy(helper[i:],nums[pr:r+1])
	copy(nums[l:r+1], helper[:])
}
