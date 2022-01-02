package sorts

import "math/rand"

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	quicksortHelper(nums, 0, len(nums)-1)
	return nums
}
func quicksortHelper(nums []int, l, r int) {
	if l >= r {
		return
	}
	// l = l + rand.Intn(r-l+1)
	l=l+int(rand.Float32()*float32(r-l+1))
	nums[l],nums[r]=nums[r],nums[l]
	p := paratition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
}
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
