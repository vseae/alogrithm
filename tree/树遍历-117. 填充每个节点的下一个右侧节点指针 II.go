package main

import "container/list"

type Node struct {
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	res := [][]*Node{}
	if root == nil { //防止为空
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []*Node
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []*Node{}        //清空层的数据
	}
	//遍历每层元素,指定next
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}
