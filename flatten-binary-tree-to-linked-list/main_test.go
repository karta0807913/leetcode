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

func flatten(current *TreeNode) {
	for current != nil {
		if current.Left != nil {
			pointer := &current.Left
			for *pointer != nil {
				pointer = &(*pointer).Right
			}
			*pointer = current.Right
			*pointer, current.Right, current.Left = current.Right, current.Left, nil
		}
		current = current.Right
	}
}

// // Morris traversal
// // this is in-order, not pre-order you idiot = =
// func flatten(current *TreeNode) {
// 	if current == nil {
// 		return
// 	}
// 	root := current
// 	current = &TreeNode{
// 		Val:   current.Val,
// 		Right: current.Right,
// 		Left:  current.Left,
// 	}
// 	hollowHead := new(TreeNode)
// 	prev := &hollowHead
// 	for current != nil {
// 		if current.Left != nil {
// 			morrisNode := &current.Left
// 			for *morrisNode != nil && (*morrisNode).Right != current {
// 				morrisNode = &(*morrisNode).Right
// 			}
// 			if *morrisNode == nil {
// 				*morrisNode = current
// 				current = current.Left
// 			} else {
// 				prev = morrisNode
// 				(*prev).Left = nil
// 				current = current.Right
// 			}
// 		} else {
// 			*prev = current
// 			prev = &current.Right
// 			if *prev != nil {
// 				(*prev).Left = nil
// 			}
// 			current = current.Right
// 		}
// 	}
// 	*root = *hollowHead
// }

func TestFlatten(t *testing.T) {
	target := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	flatten(target)
	assert.Equal(t, TreeNode{Val: 2, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}}, target)
}
