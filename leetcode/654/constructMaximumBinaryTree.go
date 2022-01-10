package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildTree(split(nums))
}

func buildTree(left, right []int, root int) *TreeNode {

	var l, r *TreeNode
	if len(left) != 0 {
		l = buildTree(split(left))
	}
	if len(right) != 0 {
		r = buildTree(split(right))
	}
	return &TreeNode{
		Val:   root,
		Left:  l,
		Right: r,
	}
}

func split(ints []int) ([]int, []int, int) {
	if len(ints) == 1 {
		return nil, nil, ints[0]
	}
	maxi := 0
	for i := 1; i < len(ints); i++ {
		if ints[maxi] < ints[i] {
			maxi = i
		}
	}
	return ints[:maxi], ints[maxi+1:], ints[maxi]
}
