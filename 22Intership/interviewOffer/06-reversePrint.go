package interviewoffer

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	cur := head
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
