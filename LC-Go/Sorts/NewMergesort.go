package Sorts

func NewmergeSort(nums []int) {
	mergeSortNew(nums, 0, len(nums)-1)

}

func mergeSortNew(nums []int, l int, r int) {
	if l >= r {
		return
	}

	mid := l + ((r - l) >> 1)
	mergeSortNew(nums, l, mid)
	mergeSortNew(nums, mid+1, r)
	mergeNew(nums, l, mid, r)
}

func mergeNew(nums []int, l int, mid int, r int) {
	p1, p2, helper, i := l, mid+1, make([]int, r-l+1), 0
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
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
	copy(nums[l:r+1], helper[:])
}
func quickSortDemo(nums []int) {
	if len(nums) == 0 {
		return
	}
	quicksortHelper(nums, 0, len(nums)-1)
}
func quicksortHelper(nums []int, l, r int) {
	if l == r {
		return
	}
	p := paration(nums, l, r)
	quickSort(nums, l, p[0]-1)
	quickSort(nums, p[1]+1, r)
}

func paration(nums []int, l, r int) []int {
	if l == r {
		return nil
	}
	less, more := l-1, r
	for l < more {
		if nums[l] > nums[r] {
			less++
			nums[less], nums[l] = nums[l], nums[less]
			l++
		} else if nums[l] < nums[r] {
			more--
			nums[more], nums[l] = nums[l], nums[more]
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}
