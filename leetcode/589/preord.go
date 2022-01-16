package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var r []int
	s(root, &r)
	return r
}

func s(root *Node, r *[]int) {
	if root == nil {
		return
	}
	*r = append(*r, root.Val)
	for _, c := range root.Children {
		s(c, r)
	}

}
