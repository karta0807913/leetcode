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

// (setq dap-print-io t)
// (require 'dap-go)
// (gethash (buffer-file-name) (dap--get-breakpoints))

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := make([]*TreeNode, 0, 4002)
	stack = append(stack, root, nil)
	currentLevel := make([]int, 0)
	result := make([][]int, 0)
	for len(stack) != 1 {
		if stack[0] == nil {
			stack = append(stack, nil)
			result = append(result, currentLevel)
			currentLevel = make([]int, 0)
			stack = stack[1:]
			continue
		}
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
		}
		currentLevel = append(currentLevel, stack[0].Val)
		stack = stack[1:]
	}
	if len(currentLevel) != 0 {
		result = append(result, currentLevel)
	}
	return result
}

func TestLevel(t *testing.T) {
	assert.Equal(t, [][]int{
		{3},
		{9, 20},
		{15, 7},
	}, levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}), nil)

	assert.Equal(t, [][]int{
		{1},
	}, levelOrder(&TreeNode{
		Val: 1,
	}), nil)
	assert.Equal(t, [][]int{}, levelOrder(nil), nil)
}
