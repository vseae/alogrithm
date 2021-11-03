package bfs

/*
	给定一个二叉树，找出其最小深度。

	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

	说明：叶子节点是指没有子节点的节点。

	输入：root = [3,9,20,null,null,15,7]
	输出：2
	示例 2：

	输入：root = [2,null,3,null,4,null,5,null,6]
	输出：5


	提示：
	树中节点数的范围在 [0, 105]内
	-1000 <= Node.val <= 1000

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree

*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func minDepth(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}

//func main() {
//
//}
