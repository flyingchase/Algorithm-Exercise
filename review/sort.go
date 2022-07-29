package review

func Heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapIfy(nums, 0, size)
	}
	return nums
}
func heapIfy(nums []int, index int, size int) {
	left := 2*index + 1
	for left < size {
		largest := left
		if left+1 < size && nums[left] < nums[left+1] {
			largest = left + 1
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = 2*index + 1
	}
}
func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] <= nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}
}
func Mergesort(nums []int) []int {
	return nums
}

func Quicksort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// sort.Ints(nums)
	quicksortHelper(nums, 0, len(nums)-1)
	return nums
}
func quicksortHelper(nums []int, l, r int) {
	if l > r {
		return
	}
	p := paratiton(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
}

func paratiton(nums []int, l, r int) []int {
	if l > r {
		return nil
	}
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

func quicksortDemo(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	return quicksortDemoHelper(nums, 0, len(nums)-1)
}
func quicksortDemoHelper(nums []int, l, r int) []int {
	if len(nums) == 0 {
		return nums
	}
	if l > r {
		return nums

	}

	p = parationDemo(nums, l, r)
	quicksortDemoHelper(nums, l, p[0]-1)
	quicksortDemoHelper(nums, p[1]+1, r)
	return nums
}

func parationDemo(nums []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if nums[l] < nums[r] {
			less++
			nums[l], nums[less] = nums[less], nums[l]
			l++
		} else if nums[l] > nums[r] {
			more--
			nums[l], nums[more] = nums[more], nums[l]
		} else {
			l++
		}
	}
	nums[r], nums[more] = nums[more], nums[r]
	return []int{less + 1, more}
}
