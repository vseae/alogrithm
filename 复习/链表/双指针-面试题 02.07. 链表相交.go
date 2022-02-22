package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curB.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curA.Next
		}
	}
	return curA
}
