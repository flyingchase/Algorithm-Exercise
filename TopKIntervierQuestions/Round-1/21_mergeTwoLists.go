package round1
func mergeTwoLists(list1 *ListNode,list2 *ListNode) *ListNode  {
	if list1==nil {
		return list2
	}
	if list2==nil {
		return list1
	}
	dummyHead:=&ListNode{Val: -1}
	n1,n2,cur:=list1,list2,dummyHead
	for n1!=nil&&n2!=nil {
		if n1.Val<=n2.Val {
			cur.Next=n1
			n1=n1.Next
		}else {
			cur.Next=n2
			n2=n2.Next
		}
		cur=cur.Next
	}
	if n1 != nil {
		cur.Next=n1
	}
	if n2!=nil {
		cur.Next=n2
	}
	return dummyHead.Next
}
