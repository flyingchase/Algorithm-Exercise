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
	copy(nums[l : r+1],helper[:])

}
