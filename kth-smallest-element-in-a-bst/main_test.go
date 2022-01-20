package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	for root != nil {
		if root.Left != nil {
			morrisPoint := &root.Left
			for (*morrisPoint) != nil && (*morrisPoint) != root {
				morrisPoint = &(*morrisPoint).Right
			}
			if (*morrisPoint) == nil {
				*morrisPoint = root
				root = root.Left
			} else {
				k--
				if k == 0 {
					return root.Val
				}
				root = root.Right
			}
		} else {
			k--
			if k == 0 {
				return root.Val
			}
			root = root.Right
		}
	}
	return 0
}
