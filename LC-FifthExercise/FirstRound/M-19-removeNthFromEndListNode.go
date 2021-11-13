package firstround

//
// import "LC-FifthExercise/ds"
//
//
// type ListNode = ds.ListNode
//
// func removeNthFromEndListNode(head *ListNode, n int) *ListNode {
// 	dummyHead := &ListNode{
// 		Next: head,
// 	}
// 	// 记录 slow 指针的上个位置，slow 和 fast 从 head 开始遍历
// 	preSlow, slow, fast := dummyHead, head, head
// 	for fast != nil {
// 		// n<=0则代表 fast 已经先走 n 步，此时 slow 开动
// 		if n <= 0 {
// 			preSlow = slow
// 			slow = slow.Next
// 		}
// 		n--
// 		fast = fast.Next
// 	}
// 	// delete slowListNode
// 	preSlow.Next = slow.Next
// 	return dummyHead.Next
// }
