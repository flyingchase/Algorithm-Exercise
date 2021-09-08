package Sorts

func NewHeapSort(nums []int) []int {
	length := len(nums)
	// construct the heap
	for i := 0; i < length; i++ {
		NewHeapInsert(nums, i)
	}
	//	sort
	for length > 0 {
		length--
		nums[length], nums[0] = nums[0], nums[length]
		heapIfy(nums, length, 0)
	}
	return nums
}

func NewHeapInsert(nums []int, index int) {
	for (index-1)/2 >= 0 && nums[index] > nums[(index-1)/2] {
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		index = (index - 1) / 2
	}

}

func newHepIfy(nums []int, size int, index int) {
	for {
		max := index
		if 2*index+1 < size && nums[2*index+1] > nums[max] {
			max = 2*index + 1
		}
		if 2*index+2 < size && nums[2*index+2] > nums[max] {
			max = 2*index + 2
		}
		if index == max {
			break
		}
		nums[index], nums[max] = nums[max], nums[index]
		index = max
	}

}
