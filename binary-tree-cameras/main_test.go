package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func helper(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, math.MaxInt32
	}
	l1, l2, l3 := helper(root.Left)
	r1, r2, r3 := helper(root.Right)

	return l2 + r2,
		min(l3+min(r2, r3), r3+min(l2, l3)),
		1 + min(l1, min(l2, l3)) + min(r1, min(r2, r3))
}

func minCameraCover(root *TreeNode) int {
	_, a, b := helper(root)
	return min(a, b)
}
