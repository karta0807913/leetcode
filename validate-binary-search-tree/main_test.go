package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	isValid := root.Right == nil || (root.Right.Val > root.Val && root.Right.Val < max)
	isValid = isValid && (root.Left == nil || (root.Left.Val < root.Val && root.Left.Val > min))
	if !isValid {
		return false
	}
	isValid = isValid && checkValidBST(root.Left, min, root.Val)
	isValid = isValid && checkValidBST(root.Right, root.Val, max)
	return isValid
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkValidBST(root, math.MinInt64, math.MaxInt64)
}
