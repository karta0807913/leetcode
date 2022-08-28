package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	next := []*TreeNode{}
	isComplete := true
	for len(queue) != 0 {
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
				if !isComplete {
					return false
				}
			}
			isComplete = node.Left != nil
			if node.Right != nil {
				next = append(next, node.Right)
				if !isComplete {
					return false
				}
			}
			isComplete = node.Right != nil
		}
		next, queue = queue, next
		next = next[:0]
	}
	return true
}
