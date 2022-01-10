package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(l, r *TreeNode) bool {
	if l == nil {
		return r == nil
	}
	if r == nil {
		return l == nil
	}
	return _isSymmetric(l.Left, r.Right) && l.Val == r.Val && _isSymmetric(l.Right, r.Left)
}
