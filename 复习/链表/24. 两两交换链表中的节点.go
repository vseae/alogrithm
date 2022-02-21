package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		//交换
		pre.Next = head.Next
		tmp := head.Next.Next
		head.Next.Next = head
		head.Next = tmp

		//更新
		pre = head
		head = tmp
	}
	return dummy.Next
}
