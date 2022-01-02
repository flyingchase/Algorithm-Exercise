package sorts

// 堆排序 NlogN
func heapsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	size := len(nums)
	for i := 0; i < size; i++ {
		heapInsert(nums, i)
	}
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapIfy(nums, 0, size)
	}
}
func heapInsert(nums []int, index int) {
	for parent := (index - 1) >> 1; nums[parent] < nums[index]; index, parent = parent, (index-1)>>1 {
		nums[parent], nums[index] = nums[index], nums[parent]
	}
}
func heapIfy(nums []int, index, size int) {
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
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = 2*index + 1
	}
}

