package Waitting

import "LC-Go/DataStructure"

type ListNode = DataStructure.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for curA != nil && curB != nil {
		lenA++
		lenB++

	}
	for curA != nil {
		lenA++
	}
	for curB != nil {
		lenB++
	}
	currA, currB := headA, headB
	if lenA > lenB {
		gap := lenA - lenB
		for gap > 0 {
			currA = currA.Next
			gap--
		}
		for currA != currB && currA != nil && currB != nil {
			currA = currA.Next
			currB = currB.Next
		}
		if currB == currA {
			return currA
		} else {
			return nil
		}

	}
	if lenA <= lenB {

		gap := lenB - lenA
		for gap > 0 {
			currB = currB.Next
			gap--
		}
		for currA != currB && currA != nil && currB != nil {
			currA = currA.Next
			currB = currB.Next
		}
		if currB == currA {
			return currA
		} else {
			return nil
		}

	}

	return nil
}
