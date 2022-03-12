package interviewoffer

import "sort"

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

// func ReversePairs(nums []int) int {
// 	if len(nums) <= 1 {
// 		return 0
// 	}
// 	count := 0
// 	mergeHelper(nums, 0, len(nums)-1, &count)
// 	return count
// }
// func mergeHelper(nums []int, l, r int, count *int) int {
// 	if l >= r {
// 		return 0
// 	}
// 	mid := l + (r-l)>>1
// 	mergeHelper(nums, l, mid, count)
// 	mergeHelper(nums, mid+1, r, count)
// 	return merge(nums, l, mid, r, count)
// }
// func merge(nums []int, left, mid, right int, count *int) int {
// 	pl, pr := left, mid+1
// 	index := 0
// 	for pl <= mid && pr <= right {
// 		if nums[pl] <= nums[pr] {
// 			pl++
// 		} else {
// 			*count++
// 			pr++
// 		}
// 		index++
// 	}
// 	return *count
// }
func ReversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tmp := append([]int{}, nums...)
	sort.Ints(tmp)
	res := 0
	for i := 0; i < len(nums); i++ {
		index := sort.SearchInts(nums, tmp[i])
		res += len(nums) - index - 1
	}
	return res
}
