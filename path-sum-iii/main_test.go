package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(node *TreeNode, targetSum int, total int, history map[int]int) int {
	if node == nil {
		return 0
	}
	result := 0
	total += node.Val
	result += history[total-targetSum]
	history[total] += 1
	result += recursive(node.Left, targetSum, total, history)
	result += recursive(node.Right, targetSum, total, history)
	history[total] -= 1
	return result
}

func pathSum(root *TreeNode, targetSum int) int {
	return recursive(root, targetSum, 0, map[int]int{0: 1})
}
