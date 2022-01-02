package sortswithgo

import "math/rand"

// 快速排序 基于分治的思想
/*
* 1. 确定分界点, 随机
* 2. 调整范围，保证 pivot为分界
* 3. 递归处理左右两段
 */
func QuickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSortHelper(nums, 0, len(nums)-1)
}
func quickSortHelper(nums []int, l, r int) {
	if l >= r {
		return
	}
	// 交换 left  right 默认选择 nums[right] 作为 pivot
	// l += int(rand.Float64() * float64(r-l+1))
	l = l + rand.Intn(r-l+1)
	nums[l], nums[r] = nums[r], nums[l]
	p := paratition(nums, l, r)
	quickSortHelper(nums, l, p[0]-1)
	quickSortHelper(nums, p[1]+1, r)
}

// 分割 less, more 代表上下界，less 和 more 之间为与 Pivot相等的值
// 返回 上下界的下标
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
