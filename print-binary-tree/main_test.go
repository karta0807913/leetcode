package main

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(current *TreeNode, height, left, right int, s [][]string) {
	if current == nil {
		return
	}
	mid := (left + right) / 2
	s[height][mid] = strconv.Itoa(current.Val)
	dfs(current.Left, height+1, left, mid, s)
	dfs(current.Right, height+1, mid+1, right, s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxHeight(current *TreeNode) int {
	if current == nil {
		return 0
	}
	return max(maxHeight(current.Left), maxHeight(current.Right)) + 1
}

func printTree(root *TreeNode) [][]string {
	height := maxHeight(root)
	width := int(math.Pow(2, float64(height))) - 1
	ans := make([][]string, height)
	for idx := range ans {
		ans[idx] = make([]string, width)
	}
	dfs(root, 0, 0, width, ans)
	return ans
}
