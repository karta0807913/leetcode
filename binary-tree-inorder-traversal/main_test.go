package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris Traversal
func inorderTraversal(current *TreeNode) (ans []int) {
	for current != nil {
		if current.Left != nil {
			morrisNode := &current.Left
			for (*morrisNode) != nil && (*morrisNode) != current {
				morrisNode = &(*morrisNode).Right
			}
			// a new traversal point
			if (*morrisNode) == nil {
				*morrisNode = current
				current = current.Left
			} else { // we already traversal left point before
				ans = append(ans, current.Val)
				current = current.Right
				*morrisNode = nil
			}
		} else if current.Right != nil {
			ans = append(ans, current.Val)
			current = current.Right
		} else {
			ans = append(ans, current.Val)
			return
		}
	}
	return
}
