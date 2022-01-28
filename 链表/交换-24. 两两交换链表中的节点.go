package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//head=list[i]
	//pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		pre.Next = next
		pre = head
		head = head.Next
	}
	return dummy.Next
}
