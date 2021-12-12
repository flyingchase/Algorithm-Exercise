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
func quicksort(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums); i++ {
	}
}
