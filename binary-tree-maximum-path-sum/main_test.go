package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recursive(root *TreeNode, maxVal *int) (ans int) {
	if root == nil {
		return 0
	}
	left := recursive(root.Left, maxVal)
	right := recursive(root.Right, maxVal)
	if left > 0 {
		ans += left
	}
	if right > 0 {
		ans += right
	}
	ans += root.Val
	if ans > *maxVal {
		*maxVal = ans
	}
	return max(root.Val+max(left, right), root.Val)
}

func maxPathSum(root *TreeNode) int {
	maxVal := math.MinInt64
	recursive(root, &maxVal)
	return maxVal
}
