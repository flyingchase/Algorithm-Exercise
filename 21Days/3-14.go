package days

// 最大子数组之和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// attention res=nums[0]
	res := nums[0]
	dp := make([]int, len(nums))
	// dp[i] means index i has maxSubArray
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func reverseLists(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 注意 prev 为 nil 避免翻转后多出一位
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummyHead.Next
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dict, res, l, r := make(map[byte]int, 0), 0, 0, 0
	for r < len(s) {
		if _, ok := dict[s[r]]; !ok {
			dict[s[r]]++
			r++
		} else {
			delete(dict, s[l])
			l++
		}
		if res < r-l {
			res = r - l
		}
	}
	return res
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	// quickSort(nums, 0, len(nums)-1)
	// mergeSort(nums, 0, len(nums)-1)
	heapSort(nums)
	return nums[len(nums)-k]
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p[0]-1)
	quickSort(nums, p[1]+1, r)
}
func partition(nums []int, l, r int) []int {
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
func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}
func merge(nums []int, l, mid, r int) {
	pl, pr := l, mid+1
	index, helper := 0, make([]int, r-l+1)
	for pl <= mid && pr <= r {
		if nums[pl] <= nums[pr] {
			helper[index] = nums[pl]
			pl++
		} else {
			helper[index] = nums[pr]
			pr++
		}
		index++
	}
	copy(helper[index:], nums[pl:mid+1])
	copy(helper[index:], nums[pr:r+1])
	copy(nums[l:r+1], helper[:])
}

func heapSort(nums []int) {
	for index := 0; index < len(nums); index++ {
		heapInsert(nums, index)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[size], nums[0] = nums[0], nums[size]
		heapIfy(nums, 0, size)
	}
}
func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	// 注意大根堆，parent<index
	for parent >= 0 && nums[parent] < nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapIfy(nums []int, index int, size int) {
	l := 2*index + 1
	for l < size {
		largest := l
		if l+1 < size && nums[l+1] > nums[l] {
			largest = l + 1
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
		l = index*2 + 1
	}
}
