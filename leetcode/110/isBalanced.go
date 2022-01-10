package _10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := height(root.Left) - height(root.Right)
	return isBalanced(root.Left) && isBalanced(root.Right) && diff <= 1 && diff >= -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := height(root.Left), height(root.Right)
	if hl > hr {
		return 1 + hl
	}
	return 1 + hr
}
