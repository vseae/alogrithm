package main

import (
	"container/list"
)

func averageOfLevels(root *TreeNode) []float64 {
	queue := list.New()
	var res [][]int
	var tmp []int
	var ans []float64
	if root == nil {
		return []float64{}
	}
	queue.PushBack(root)
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
		sum := 0
		count := 0
		for j := 0; j < len(res[i]); j++ {
			sum += res[i][j]
		}
		avg := float64(sum) / float64(count)

		ans = append(ans, avg)
	}
	return ans
}
