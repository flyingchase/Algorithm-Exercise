package intership

func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return quicksortHelper(nums, 0, len(nums)-1)
}
func quicksortHelper(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}
	mid := l + (r-l)>>1
	quicksortHelper(nums, l, mid)
	quicksortHelper(nums, mid+1, r)

	return quicksort
}
