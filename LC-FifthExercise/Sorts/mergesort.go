package sorts

func mergesort(nums []int) []int {
	if len(nums) == 0 {
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

	p1, p2, i, helper := l, mid+1, 0, make([]int, r-l+1)
	for p1 <= mid && p2 <= r {
		if nums[p1] < nums[p2] {
			helper[i] = nums[p1]
			p1++
		} else {
			helper[i] = nums[p2]
			p2++
		}
		i++
	}
	copy(helper[i:], nums[p1:mid+1])
	copy(helper[i:], nums[p2:r+1])
	copy(nums, helper)
}
