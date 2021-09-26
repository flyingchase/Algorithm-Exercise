package main

import (
	"LC-Go/DataStructure"
	"fmt"
)

type TreeNode = DataStructure.TreeNode

//func maxNodeInLevel(root *TreeNode) int  {
//if root==nil {
//	return 0
//}
//cur:=root
//queue:=make([][]int,0)
//
//for cur!=nil&&len(queue)>0 {
//	for cur!=nil {
//		size:=len(queue)
//		lists:=make([]int,0)
//		for size>0 {
//			node:=queue[size-1]
//			if node==nil {
//
//			}
//
//		}
//
//
//	}
//

//}
func main() {
	nums := []int{1, 5, 9, 8, 3, 2, 0, 4, 7, 6, 11}
	heapsort(nums)
	fmt.Println(nums)
}
func heapsort(nums []int) {
	if nums == nil {
		return
	}
	for i, _ := range nums {
		heapifynow(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapifytrue(nums, 0, size)
	}
}

func heapifytrue(nums []int, index int, size int) {
	left := index*2 + 1
	largest := left
	for left < size {
		largest = left
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		}
		if nums[index] > nums[largest] {
			index = largest
			break
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
		left = 2*index + 1

	}

}

func heapifynow(nums []int, index int) {
	parent := (index - 1) / 2
	for parent < len(nums) && parent >= 0 && nums[parent] < nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) / 2
	}

}
func mergesort(nums []int) {
	if nums == nil {
		return
	}
	mergesorthelper(nums, 0, len(nums)-1)
}
func mergesorthelper(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergesorthelper(nums, mid+1, r)
	mergesorthelper(nums, l, mid)
	merge(nums, l, mid, r)
}

func merge(nums []int, l int, mid int, r int) {
	p1, p2, i, helper := l, mid+1, 0, make([]int, 1+r-l)
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			helper[i] = nums[p1]
			p1++
		} else {
			helper[i] = nums[p2]
			p2++
		}
		i++
	}
	copy(helper[i:], nums[p1:])
	copy(helper, nums[p2:])
	copy(nums[l:r+1], helper)
}

func quicksortnow (nums[] int)   {
	if len(nums)==0 {
		return
	}
	quicksorthelper(nums)
}