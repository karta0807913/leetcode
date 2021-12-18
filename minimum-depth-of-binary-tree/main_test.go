package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root = nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root, nil)
	currentLevel := 1
	for len(stack) != 1 {
		if stack[0] == nil {
			stack = append(stack, nil)
			currentLevel += 1
			stack = stack[1:]
			continue
		}
		fmt.Println(stack[0].Val)
		flag := true
		if stack[0].Left != nil {
			flag = false
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			flag = false
			stack = append(stack, stack[0].Right)
		}
		if flag {
			return currentLevel
		}
		stack = stack[1:]
	}
	return currentLevel - 1
}

func TestDepth(t *testing.T) {
	assert.Equal(t, 2, minDepth(&TreeNode{
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

	assert.Equal(t, 3, minDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
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
}
