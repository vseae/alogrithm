package main

import (
	"container/list"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var res [][]int
	var ans []int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	for queue.Len() > 0 {
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
		tmp = []int{}
	}
	for i := 0; i < len(res); i++ {
		sort.Ints(res[i])
		ans = append(ans, res[i][len(res[i])-1])
	}
}
