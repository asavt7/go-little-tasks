package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)

	return root
}
