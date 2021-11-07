package offerInterview

func heapSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	for index := 0; index < len(nums); index++ {
		heapSortInsert(nums, index)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[size], nums[0] = nums[0], nums[size]
		heapSortIfy(nums, 0, size)
	}
	return nums
}
func heapSortInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] < nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapSortIfy(nums []int, index, size int) {
	left := 2*index + 1
	largest := left
	right := 2*index + 2
	for left < size {
		if right < size && nums[left] < nums[right] {
			largest = right
		}
		if nums[largest] < nums[index] {
			largest = index
			break
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
		left = index*2 + 1
		right = 2*index + 2
	}
}
