package Waitting

import "LC-Go/DataStructure"

type ListNode = DataStructure.ListNode
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA,curB:=headA,headB
	lenA,lenB:=0,0
	for curA!=nil {
		lenA++
	}
	for curB != nil {
	    lenB++
	}
	currA,currB:=headA, headB
	if lenA>lenB {
		gap:=lenA-lenB
		for  gap > 0{
			currA=currA.Next
			gap--
		}
		for currA!=currB {
			currA=currA.Next
			currB=currB.Next
		}

	}

}