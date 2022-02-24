package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, 9, 2, 3, 8, 7, 4, 6, 5}
	fmt.Println("before sort the array is ", nums)
	heapsort(nums)
	fmt.Println("after sort the array is ", nums)
}
func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return quicksortHelper(nums, 0, len(nums)-1)
}
func quicksortHelper(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}
	p := paratition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
	return nums
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
func mergesort(nums []int) {

	if len(nums) <= 1 {
		return
	}
	mergeHelp(nums, 0, len(nums)-1)
}
func mergeHelp(nums []int, l, r int) {
	if l >= r {
		return
	}
	if l < r {
		mid := l + (r-l)>>1
		mergeHelp(nums, l, mid)
		mergeHelp(nums, mid+1, r)
		merge(nums, l, mid, r)
	}

}
func merge(nums []int, l, mid, r int) {
	pl, pr, helper := l, mid+1, make([]int, r-l+1)
	i := 0
	for pl <= mid && pr <= r {
		if nums[pl] <= nums[pr] {
			helper[i] = nums[pl]
			pl++
		} else {
			helper[i] = nums[pr]
			pr++
		}
		i++
	}
	copy(helper[i:], nums[pl:mid+1])
	copy(helper[i:], nums[pr:r+1])
	// 注意选择 l:r+1范围内
	copy(nums[l:r+1], helper)
}

func heapsort(nums []int) {
	if len(nums) <= 1 {
		return
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
}
func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] <= nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}

}
func heapIfy(nums []int, index, size int) {
	left := index*2 + 1
	for left < size {
		largest := left
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		}
		if nums[index] > nums[largest] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = (index * 2) + 1
	}
}
