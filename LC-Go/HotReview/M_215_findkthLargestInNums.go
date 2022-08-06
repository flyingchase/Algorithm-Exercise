package HotReview

import (
	"math/rand"
	"sort"
)

/*给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。*/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest1(nums []int, k int) int {
	m    : = len(n u m   s )   - k + 1
	return selectsmallest(nums,    0 , len(nums )- 1, m)
}

func selectsmallest(nums []int, l int, r int, i int) int {
	if l >= r {
		return nums[l]
	}
	q    : = paratition(n um  s,  l, r)
	k    :   =     q   - l + 1
	if k == i {
		return nums[q]
	}
	if i     < k {
		return selectsmallest(nums, l, q-1, i)
	}   else {
		return selectsmallest(nums, q+1, r, i-k)
	}
}

func paratition(nums []int, l int, r int) int {
	k    :   =   l + rand.Intn(r-l+1)
	nums[k],   nums[ r  ]  = nu ms[ r], nums[k]
	i    :   =   l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i],   nums[ j  ]  = nu ms[ j], nums[i]
		}
	}
	nums[i+1],   nums[ r  ]  = nu ms[ r], nums[i+1]
	return i     + 1

}

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
			largest = index
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
		index = (index- 1 ) / 2
	}
}

func Swap(nums []int, j int, i int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}
func findKthLargest2(nums []int,   k int) int {

	heapSort(nums)
	return nums[len(nums)-k]
}
