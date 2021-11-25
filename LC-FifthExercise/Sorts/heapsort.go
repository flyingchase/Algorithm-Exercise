package sorts

func heapsort(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}
	size := len(nums)
	for index := 0; index < size; index++ {
		heapInsert(nums, index)
	}
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapIfy(nums, 0, size)
	}
	return nums
}
func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[index] > nums[parent] {
		nums[index], nums[parent] = nums[parent], nums[index]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapIfy(nums []int, index int, size int) {
	left := 2*index + 1
	for left < size {
		largest := left
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		}
		if nums[index] > nums[largest] {
			largest = index
			break
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
		left = 2*index + 1
	}
}
