package offerInterview

func quickSort(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}
	quickSortHelper(nums, 0, len(nums)-1)
	return nums
}

func quickSortHelper(nums []int, l, r int) {
	if l < r {
		nums[l], nums[r] = nums[r], nums[l]
		p := paratition(nums, l, r)
		quickSortHelper(nums, l, p[0]-1)
		quickSortHelper(nums, p[1]+1, r)
	}
}
func paratition(nums []int, l, r int) []int {

	less, more := l-1, r
	for l < more {

		if nums[l] < nums[r] {
			less++
			nums[less], nums[l] = nums[l], nums[less]
			l++
		} else if nums[l] > nums[r] {
			more--
			nums[more], nums[l] = nums[l], nums[more]
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}
