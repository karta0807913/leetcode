package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSorting(root *TreeNode, prev, first, second **TreeNode) {
	if root.Val < (*prev).Val {
		if *first == nil {
			*first = *prev
		}
		*second = root
	}
	*prev = root
}

// Morris Traversal
func recoverTree(root *TreeNode) {
	var first, second *TreeNode
	prev := &TreeNode{Val: math.MinInt64}
	for root != nil {
		if root.Left != nil {
			pointer := &root.Left
			for (*pointer) != nil && (*pointer) != root {
				pointer = &(*pointer).Right
			}
			if (*pointer) == nil {
				(*pointer) = root
				root = root.Left
			} else {
				(*pointer) = nil
				checkSorting(root, &prev, &first, &second)
				root = root.Right
			}
		} else {
			checkSorting(root, &prev, &first, &second)
			root = root.Right
		}
	}

	first.Val, second.Val = second.Val, first.Val
}

// // using iterate
// func iterate(root *TreeNode, prev, first, second **TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	fmt.Println((*prev).Val)
// 	iterate(root.Left, prev, first, second)
// 	if (*prev).Val > root.Val {
// 		if *first != nil {
// 			*second = root
// 		} else {
// 			*first = *prev
// 			*second = root
// 		}
// 	}
// 	*prev = root
// 	iterate(root.Right, prev, first, second)
// }

// func recoverTree(root *TreeNode) *TreeNode {
// 	prev := &TreeNode{Val: math.MinInt64}
// 	first, second := new(*TreeNode), new(*TreeNode)

// 	iterate(root, &prev, first, second)
// 	(*first).Val, (*second).Val = (*second).Val, (*first).Val
// 	return root
// }

// func TestRecover(t *testing.T) {
// 	assert.Equal(t, &TreeNode{
// 		Val: 3,
// 		Left: &TreeNode{
// 			Val:   1,
// 			Right: &TreeNode{Val: 2},
// 		},
// 	}, recoverTree(&TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val:   3,
// 			Right: &TreeNode{Val: 2},
// 		},
// 	}))
// }
