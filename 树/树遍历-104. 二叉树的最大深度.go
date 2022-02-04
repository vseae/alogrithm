package main

import "container/list"

/**
 * Definition for a binary 树 node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	queue := list.New()
	if root == nil {
		return 0
	}
	var res [][]int
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := []int{}
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return len(res)
}
