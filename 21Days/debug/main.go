package main

import (
	days "21days"
)

type ListNode = days.ListNode

func main() {
	l1 := &ListNode{1, &ListNode{2, nil}}
	// l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// days.SortList(head)
	// days.ReorderList(head)
	// for head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// }
}
