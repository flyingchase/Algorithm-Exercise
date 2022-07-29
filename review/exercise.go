package review

func heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	return nums
}
func heapinsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] < nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapify(nums []int, index, size int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		largest := left
		if right < size && nums[right] > nums[left] {
			largest = right
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left, right = 2*index+1, 2*index+2
	}
}

func quicksort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	quicsortHelper2(nums, 0, len(nums)-1)
	return nums
}
func quicsortHelper2(nums []int, l, r int) {
	if l > r {
		return
	}
	p := paratiton2(nums, l, r)
	quicsortHelper2(nums, l, p[0]-1)
	quicsortHelper2(nums, p[1]+1, r)
}
func paratiton2(nums []int, l, r int) []int {
	if l > r {
		return []int{}
	}
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
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}
