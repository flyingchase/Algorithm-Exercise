package exercise

func heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	size := len(nums)
	for i := 0; i < size; i++ {
		heapinsert(nums, i)
	}
	for size > 0 {
		size--
		nums[size], nums[0] = nums[0], nums[size]
		heapIfy(nums, size, 0)
	}
	return nums
}
func heapIfy(nums []int, size, index int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		largest := left
		if right < size && nums[right] < nums[left] {
			largest = right
		}
		if nums[index] < nums[largest] {
			largest = index
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left, right = 2*index+1, 2*index+2
	}
}
func heapinsert(nums []int, index int) {
	parent := (index - 1) / 2
	for parent >= 0 && nums[index] < nums[parent] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) / 2
	}
}
func quicksort(nums []int) {
	if len(nums) == 0 {

	}

}
