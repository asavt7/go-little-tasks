package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || _isSubtree(root, subRoot) || isSubtree(root.Right, subRoot)
}

func _isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return root == nil
	}

	return _isSubtree(root.Left, subRoot.Left) && _isSubtree(root.Right, subRoot.Right) && root.Val == subRoot.Val
}
