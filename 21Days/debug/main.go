package main

import (
	days "21days"
	"fmt"
)

type ListNode = days.ListNode

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	days.ReorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
