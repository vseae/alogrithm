package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}
	return true
}
