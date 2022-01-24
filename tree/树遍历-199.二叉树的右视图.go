package main

import "container/list"

/*
输入: [1,2,3,null,5,null,4]
输出: [1,3,4]
*/
func rightSideView(root *TreeNode) []int {
	queue := list.New()
	var res [][]int
	var ans []int
	if root == nil {
		return []int{}
	}
	var tmp []int
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
		ans = append(ans, res[i][len(res[i])-1])
	}
	return ans
}
