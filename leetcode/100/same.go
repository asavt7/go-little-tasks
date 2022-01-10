package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return q == p
	}

	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
}

func main() {
	var a *TreeNode
	var b *TreeNode = &TreeNode{}
	println(a == b)
}
