package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNext(node *TreeNode) int {
	sum := 0
	if node == nil {
		return sum
	}
	if node.Left != nil {
		sum += node.Left.Val
	}
	if node.Right != nil {
		sum += node.Right.Val
	}
	return sum
}

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val&1 == 0 {
		sum += sumNext(root.Left)
		sum += sumNext(root.Right)
	}
	sum += sumEvenGrandparent(root.Left)
	sum += sumEvenGrandparent(root.Right)
	return sum
}

func TestSum(t *testing.T) {
}
