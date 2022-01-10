package paths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	s := strconv.Itoa(root.Val)
	res := make([]string, 0)
	foo := func(ss string) {
		res = append(res, ss)
	}
	if root.Left == nil && root.Right == nil {
		foo(s)
	} else {
		bst(root.Left, s, foo)
		bst(root.Right, s, foo)
	}
	return res
}

func bst(root *TreeNode, s string, f func(ss string)) {
	if root == nil {
		return
	}
	s += "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		f(s)
		return
	}
	bst(root.Left, s, f)
	bst(root.Right, s, f)
}
