package datawhale50

import (
	"sort"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	// 避免循环 k 移动
	k %= length
	for i := 0; i < k; i++ {
		head = rotateRightOne(head)
	}
	return head
}
func rotateRightOne(head *ListNode) *ListNode {
	prev, cur := head, head
	// 找到当前链尾
	for cur.Next != nil {
		prev, cur = cur, cur.Next
	}
	cur.Next = head
	head = cur
	prev.Next = nil
	return head
}
func merge(nums1 []int, nums2 []int, m, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	// quickSort(nums, 0, len(nums)-1)

	length := len(nums)
	for index := 0; index < length; index++ {
		heapinsert(nums, index)
	}
	for length > 0 {
		length--
		nums[length], nums[0] = nums[0], nums[length]
		heapify(nums, 0, length)
	}
	return nums
}
func heapinsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[index] < nums[parent] {
		nums[index], nums[parent] = nums[parent], nums[index]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapify(nums []int, index int, size int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		largest := left
		if nums[left] <= nums[right] {
			largest = right
		}
		if nums[index] >= nums[largest] {
			largest = index
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = 2*index + 1
		right = 2*index + 2
	}
}
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := paratition(nums, l, r)
	quickSort(nums, l, p[0]-1)
	quickSort(nums, p[1]+1, r)
}
func paratition(nums []int, l, r int) []int {
	if l >= r {
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
