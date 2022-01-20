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

const b = 1 << 61

func updateNode(n int, node *TreeNode) {
	if node != nil && node.Val&b == 0 {
		node.Val = node.Val + (n&(^b))*10
		node.Val |= b
	}
}

func sumNumbers(current *TreeNode) int {
	if current != nil {
		current.Val |= b
	}
	sum := 0
	for current != nil {
		if current.Left != nil {
			morrisPoint := current.Left
			for morrisPoint.Right != nil && morrisPoint.Right != current {
				morrisPoint = morrisPoint.Right
			}
			if morrisPoint.Right == nil {
				morrisPoint.Right = current
				updateNode(current.Val, current.Left)
				current = current.Left
			} else {
				morrisPoint.Right = nil
				if morrisPoint.Left == nil {
					sum += morrisPoint.Val & (^b)
					// fmt.Println(morrisPoint.Val& (^b))
				}
				updateNode(current.Val, current.Right)
				current = current.Right
			}
		} else if current.Right != nil {
			updateNode(current.Val, current.Right)
			current = current.Right
		} else {
			sum += current.Val & (^b)
			// fmt.Println(current.Val& (^b))
			current = nil
		}
	}
	// fmt.Println()
	return sum
}

func TestNumber(t *testing.T) {
	assert.Equal(t, 25, sumNumbers(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}
