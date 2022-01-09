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
			newRoot = root
			newRoot.Left = nil
			cur = newRoot
		} else {
			cur.Right = root
			cur = cur.Right
			cur.Left = nil
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
