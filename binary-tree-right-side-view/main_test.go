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

// so sad, the channel version is slower :'(
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 2, 10)
	queue[0] = root
	queue[1] = nil
	var prev, node *TreeNode
	for {
		queue, node = queue[1:], queue[0]
		if node == nil {
			if prev == nil {
				return
			}
			ans = append(ans, prev.Val)
			queue = append(queue, nil)
			prev = nil
		} else {
			prev = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}

/*
// channel version
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	// haha
	queue := make(chan *TreeNode, 64)
	queue <- root
	queue <- nil

	var prev *TreeNode
	for node := range queue {
		if node == nil {
			if prev == nil {
				return
			}
			ans = append(ans, prev.Val)
			queue <- nil
			prev = nil
		} else {
			prev = node
			if node.Left != nil {
				queue <- node.Left
			}
			if node.Right != nil {
				queue <- node.Right
			}
		}
	}
	return
}
*/

func TestView(t *testing.T) {
	assert.Equal(t, []int{1, 3, 4}, rightSideView(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}))
}
