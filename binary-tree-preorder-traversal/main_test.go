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

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		if queue[0] == nil {
			queue = queue[1:]
			continue
		}
		result = append(result, queue[0].Val)
		queue = append([]*TreeNode{queue[0].Left, queue[0].Right}, queue[1:]...)
	}
	return result
}

func TestPreorder(t *testing.T) {
	assert.Equal(t, []int{1, 4, 3, 2}, preorderTraversal(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}), nil)
}
