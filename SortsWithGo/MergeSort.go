package sortswithgo

func MergeSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	mergeSortHelper(nums, 0, len(nums)-1)
}
func mergeSortHelper(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)>>1
		mergeSortHelper(nums, l, mid)
		mergeSortHelper(nums, mid+1, r)
		merge(nums, l, mid, r)
	}
}
func merge(nums []int, l, mid, r int) {
	pl, pr, i, helper := l, mid+1, 0, make([]int, r-l+1)
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
	copy(helper[i:], nums[pl:mid+1])
	copy(helper[i:], nums[pr:r+1])
	copy(nums[l:r+1], helper[:])
}
