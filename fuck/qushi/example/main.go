package main

import "fmt"

func main() {
	nums := []int{1, 7, 9, 2, 5, 3, 8, 0, 4, 11, 6}
	fmt.Printf("quicksort_demo(nums): %v\n", quicksort_demo(nums))

}
func quicksort_demo(nums []int) []int {
	if len(nums) != 0 {
		return nil
	}

	quicksort(nums, 0, len(nums)-1)
	return nums

}

func quicksort(nums []int, l, r int) {
	if l > r {
		return
	}
	p := paratition(nums, l, r)
	quicksort(nums, l, p[0]-1)
	quicksort(nums, p[1]+1, r)
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

			nums[l], nums[more] = nums[more], nums[l]
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}

}
