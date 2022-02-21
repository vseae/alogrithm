package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	ans := robTree(root)
	return max(ans[0], ans[1])
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	left := robTree(cur.Left)
	right := robTree(cur.Right)
	robCur := cur.Val + left[0] + right[0]
	noRobCur := max(left[0], left[1]) + max(right[0], right[1])
	return []int{noRobCur, robCur}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
