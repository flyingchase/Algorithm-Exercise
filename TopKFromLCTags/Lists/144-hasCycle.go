package topkLists

// two pointers fast go two steps and slow go one steps
// remember to adjust
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head.Next
	for fast != slow {
		if fast.Next == nil || fast == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// set implmented by struct{}
// 空结构体占位
func hasCycle2(head *ListNode) bool {
	dict := make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := dict[cur]; ok {
			return true
		}
		dict[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}
func Quicksort(nums []int) {
	if len(nums) == 0 {
		return
	}
	quicksortHelper(nums, 0, len(nums)-1)
}
func quicksortHelper(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)

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
	nums[r], nums[more] = nums[more], nums[r]
	return []int{less + 1, more}
}
