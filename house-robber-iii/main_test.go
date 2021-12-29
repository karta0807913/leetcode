package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func robNode(root *TreeNode) (prev, pprev, ppprev int) {
	if root == nil {
		return
	}
	prevLeft, pprevLeft, ppprevLeft := robNode(root.Left)
	prevRight, pprevRight, ppprevRight := robNode(root.Right)
	prev = max(pprevRight, ppprevRight) + max(pprevLeft, ppprevLeft) + root.Val
	prev = max(prev, max(prevLeft, pprevLeft)+max(prevRight, pprevRight))
	pprev = prevLeft + prevRight
	ppprev = pprevRight + pprevLeft
	// fmt.Printf("prev: %d, pprev: %d, ppprev: %d\n", prev, pprev, ppprev)
	return
}

func rob(root *TreeNode) int {
	prev, pprev, _ := robNode(root)
	return max(prev, pprev)
}

func TestRob(t *testing.T) {
	assert.Equal(t, 7, rob(
		&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	))
	assert.Equal(t, 7, rob(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 1},
			},
		},
	))

	assert.Equal(t, 9, rob(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val: 5,
				// Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 1},
			},
		},
	))

	assert.Equal(t, 9, rob(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val: 1,
				// Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
		},
	))

	assert.Equal(t, 11, rob(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val: 1,
				// Left:  &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 1},
				},
			},
		},
	))

	assert.Equal(t, 15, rob(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 1},
			},
		},
	))
}
