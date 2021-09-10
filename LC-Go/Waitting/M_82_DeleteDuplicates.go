package Waitting


func deleteDuplicates(head *ListNode) *ListNode {
	m:=make(map[int]int ,1)
	cur:=head
	for cur!=nil {
		m[cur.Val]++
		cur=cur.Next
	}

	dummyHead:=&ListNode{Val: -1,Next: head}
	newCur:=dummyHead.Next
	prev:=dummyHead
	for newCur!=nil {
		if m[newCur.Val]>1 {
			prev,newCur=deleteNode(prev,newCur)
			continue
		}
		newCur=newCur.Next
		prev=prev.Next
	}
	return dummyHead.Next
}

func deleteNode(prev,node *ListNode)  (*ListNode,*ListNode) {
	prev.Next=node.Next
	node=node.Next
	return prev,node
}