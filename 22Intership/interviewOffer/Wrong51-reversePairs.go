package interviewoffer

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
func ReversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := 0
	res = mergePairs(nums, 0, len(nums)-1)
	return res
}
func mergePairs(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := l + (r-l)>>1
	mergePairs(nums, l, mid)
	mergePairs(nums, mid+1, r)
	return merge(nums, l, mid, r)
}
func merge(nums []int, l, mid, r int) int {
	pl, pr := l, mid+1
	res := 0
	for pl <= mid && pr <= r {
		if nums[pl] < nums[pr] {
			res++
			pl++
		} else if nums[pl] == nums[pr] {
			pl++
			pr++
		} else {
			pr++
		}
	}
	for pl <= mid && nums[pl] < nums[r] {
		res++
		pl++
	}
	return res
}
