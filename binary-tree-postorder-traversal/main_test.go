package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(current *TreeNode) (ans []int) {
	stack := make([]*TreeNode, 1)
	stack[0] = current
	prev := &TreeNode{}
	for len(stack) != 0 {
		current = stack[len(stack)-1]
		if prev == current.Right {
			ans = append(ans, current.Val)
			stack = stack[:len(stack)-1]
		} else if current.Left != nil && prev != current.Left {
			stack = append(stack, current.Left)
		} else if current.Right != nil {
			stack = append(stack, current.Right)
		} else {
			ans = append(ans, current.Val)
			stack = stack[:len(stack)-1]
		}
		prev = current
	}
	return
}
