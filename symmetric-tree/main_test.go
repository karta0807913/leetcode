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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftNodes := []*TreeNode{root.Left}
	rightNodes := []*TreeNode{root.Right}
	for len(leftNodes) != 0 && len(rightNodes) != 0 {
		if leftNodes[0] == nil && rightNodes[0] == nil {
			leftNodes = leftNodes[1:]
			rightNodes = rightNodes[1:]
			continue
		}
		if (leftNodes[0] != rightNodes[0]) && (leftNodes[0] == nil || rightNodes[0] == nil) {
			return false
		}
		if leftNodes[0].Val != rightNodes[0].Val {
			return false
		}
		leftNodes = append(leftNodes, leftNodes[0].Left, leftNodes[0].Right)
		rightNodes = append(rightNodes, rightNodes[0].Right, rightNodes[0].Left)
		leftNodes = leftNodes[1:]
		rightNodes = rightNodes[1:]
	}

	return len(leftNodes) == len(rightNodes)
}

func TestMethod(t *testing.T) {
	assert.True(t, isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
			Left: &TreeNode{
				Val: 4,
			},
		},
	}), nil)

	assert.True(t, isSymmetric(&TreeNode{
		Val: 1,
	}), nil)

	assert.False(t, isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
			Left: &TreeNode{
				Val: 4,
			},
		},
	}), nil)

	assert.False(t, isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
			Left: &TreeNode{
				Val: 4,
			},
		},
	}), nil)
}
