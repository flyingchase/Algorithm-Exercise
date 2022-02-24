package Sorts

func heapSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		Swap(nums, 0, size)
		heapIfy(nums, 0, size)
	}
	return nums
}
func heapIfy(nums []int, index int, size int) {
	left := index*2 + 1
	for left < size {
		var largest int
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if nums[largest] < nums[index] {
			break
		}
		Swap(nums, index, largest)
		index = largest
		left = 2*index + 1
	}
}

func heapInsert(nums []int, index int) {
	for nums[index] > nums[(index-1)/2] {
		Swap(nums, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func Swap(nums []int, j int, i int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}
