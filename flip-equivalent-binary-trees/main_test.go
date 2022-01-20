package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func EqualNode(left, right *TreeNode) bool {
	if left == right && left == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if !EqualNode(root1, root2) {
		return false
	}
	if root1 == nil {
		return true
	}

	if EqualNode(root1.Left, root2.Left) {
		return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
	}
	return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}
