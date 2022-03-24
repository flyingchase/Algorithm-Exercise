package main

import (
	days "21days"
	"fmt"
)

type ListNode = days.ListNode

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// days.SortList(head)
	// days.ReorderList(head)
	head := days.AddTwoNumbers(l1, l2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
