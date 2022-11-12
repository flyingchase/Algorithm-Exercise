package tricks

import (
	"math/rand"
	"time"
)

func Quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	seed := time.Now().Unix()
	rand.Seed(seed)
	pivot := rand.Intn(r - l)
	pivot += l
	nums[r], nums[pivot] = nums[pivot], nums[r]

	p := paratitoin(nums, l, r)

	Quicksort(nums, l, p[0]-1)
	Quicksort(nums, p[1]+1, r)
}

func paratitoin(nums []int, l, r int) []int {

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

func Mergesort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)>>1
	Mergesort(nums, l, mid)
	Mergesort(nums, mid+1, r)
	mergeSortHelper(nums, l, mid, r)
}

func mergeSortHelper(nums []int, l, mid, r int) {
	helper := make([]int, r-l+1)
	p1, p2, index := l, mid+1, 0
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			helper[index] = nums[p1]
			p1++
		} else {
			helper[index] = nums[p2]
			p2++
		}
		index++
	}

	copy(helper[index:], nums[p1:mid+1])
	copy(helper[index:], nums[p2:r+1])
	copy(nums[l:r+1], helper[:])
}

func Heapsort(nums []int) {
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}

	size := len(nums)
	for size > 0 {
		size--
		nums[size], nums[0] = nums[0], nums[size]
		heapIfy(nums, 0, size)
	}

}

func heapIfy(nums []int, index, size int) {
	left, right := index*2+1, index*2
	for left < size {
		largest := left
		if right < size && nums[right] > nums[left] {
			largest = right
		}
		if nums[index] > nums[largest] {
			largest = index
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left, right = index*2+1, index*2
	}

}

func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] < nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}
}
