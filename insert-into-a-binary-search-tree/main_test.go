package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	pointer := &root
	for *pointer != nil {
		if val < (*pointer).Val {
			pointer = &(*pointer).Left
		} else {
			pointer = &(*pointer).Right
		}
	}
	*pointer = &TreeNode{
		Val: val,
	}
	return root
}
