package main

/**
 * Definition for a binary 树 node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var vals []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	l, r := 0, len(vals)-1
	sum := 0
	for l < r {
		sum = vals[l] + vals[r]
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return false
}

func inOrder(node *TreeNode, vals []int) {
	if node == nil {
		return
	}
	inOrder(node.Left, vals)
	vals = append(vals, node.Val)
	inOrder(node.Right, vals)
}
