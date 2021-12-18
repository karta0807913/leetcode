package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root, nil)
	currentLevel := make([]int, 0)
	result := make([][]int, 0)
	for len(stack) != 1 {
		if stack[0] == nil {
			stack = append(stack, nil)
			result = append(result, currentLevel)
			currentLevel = make([]int, 0)
			stack = stack[1:]
			continue
		}
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
		}
		currentLevel = append(currentLevel, stack[0].Val)
		stack = stack[1:]
	}
	if len(currentLevel) != 0 {
		result = append(result, currentLevel)
	}
	for idx := 0; idx < len(result)/2; idx++ {
		result[idx], result[len(result)-1-idx] = result[len(result)-1-idx], result[idx]
	}
	return result
}
