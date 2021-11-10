package main

/**
 * Definition for a binary tree node.
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

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	//中序遍历
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			vals = append(vals, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, val := range vals {
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return dummyNode.Right
}

func main() {

}
