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

func find(arr []int, target int) int {
	for idx := range arr {
		if arr[idx] == target {
			return idx
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var splitPoint, leaderPoint int
	for leaderPoint = range preorder {
		splitPoint = find(inorder, preorder[leaderPoint])
		if splitPoint != -1 {
			break
		}
	}
	return &TreeNode{
		Val:   preorder[leaderPoint],
		Left:  buildTree(preorder[leaderPoint+1:], inorder[:splitPoint]),
		Right: buildTree(preorder[leaderPoint+1:], inorder[splitPoint+1:]),
	}
}

func TestBuildTree(t *testing.T) {
	assert.Equal(t, &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}, buildTree([]int{
		3, 9, 20, 15, 7,
	}, []int{
		9, 3, 15, 20, 7,
	}))
}
