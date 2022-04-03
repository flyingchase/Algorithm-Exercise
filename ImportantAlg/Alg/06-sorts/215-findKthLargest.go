package sorts

// M-215-数组中的第 k 大元素
func findKthLargest(nums []int, k int) int {

	return nums[len(nums)-k+1]
}
