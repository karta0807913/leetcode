package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	current := []*TreeNode{root}
	next := make([]*TreeNode, 0, 2)
	answer := map[int][]int{
		0: []int{root.Val},
	}
	root.Val = 0
	for len(current) != 0 {
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
				nextKey := node.Val - 1
				answer[nextKey] = append(answer[nextKey], node.Left.Val)
				node.Left.Val = nextKey
			}
			if node.Right != nil {
				next = append(next, node.Right)
				nextKey := node.Val + 1
				answer[nextKey] = append(answer[nextKey], node.Right.Val)
				node.Right.Val = nextKey
			}
		}
		next, current = current, next
		next = next[:0]
	}
	result := make([][]int, 0, len(answer))
	for i := 0; i <= 1000; i++ {
		if len(answer[i]) != 0 {
			result = append(result, answer[i])
		}
	}
	return result
}
