package Sorts

func MergeSort(nums []int) []int {
	if nums == nil {
		return nums
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + ((r - l) >> 1)
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	// 传递引用 此处不是内部递归调用 调用 merge
	merge(&nums, l, mid, r)

}

func merge(nums *[]int, l int, mid int, r int) {

	p1, p2 := l, mid+1
	helper := make([]int, r-l+1)
	i := 0
	for p1 <= mid && p2 <= r {
		if (*nums)[p1] <= (*nums)[p2] {
			helper = append(helper, (*nums)[p1])
			p1++
		} else {
			helper = append(helper, (*nums)[p2])
			p2++
		}
		i++
	}
	for j := 0; j < i; j++ {
		(*nums)[j+l] = helper[j]
	}
}
