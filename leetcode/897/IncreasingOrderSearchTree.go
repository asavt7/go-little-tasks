package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var newRoot *TreeNode
	var cur *TreeNode

	foo := func(root *TreeNode) {
		if newRoot == nil {
			newRoot = &TreeNode{
				Val: root.Val,
			}
			cur = newRoot
		} else {
			cur.Right = &TreeNode{
				Val: root.Val,
			}
			cur = cur.Right
		}
	}
	bst(root, foo)
	return newRoot
}

func bst(root *TreeNode, foo func(root *TreeNode)) {
	if root == nil {
		return
	}
	bst(root.Left, foo)
	foo(root)
	bst(root.Right, foo)
}

func createNewtree() {

}
