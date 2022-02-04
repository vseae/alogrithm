package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

/*
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
*/
func levelOrder(root *Node) [][]int {
	queue := list.New()
	res := [][]int{} //结果集
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len() //记录当前层的数量
		var tmp []int
		for T := 0; T < length; T++ { //该层的每个元素：一添加到该层的结果集中；二找到该元素的下层元素加入到队列中，方便下次使用
			myNode := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res, tmp)
	}
	return res
}
