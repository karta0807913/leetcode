package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	prevNode := []*TreeNode{root}
	nextNode := []*TreeNode{}
	result := 0
	for len(prevNode) != 0 {
		for _, node := range prevNode {
			if node.Left != nil {
				nextNode = append(nextNode, node.Left)
			}
			if node.Right != nil {
				nextNode = append(nextNode, node.Right)
			}
		}
		// select sort
		for i := 0; i < len(nextNode); i++ {
			var minNodeIndex = i
			for j := i + 1; j < len(nextNode); j++ {
				if nextNode[j].Val < nextNode[minNodeIndex].Val {
					minNodeIndex = j
				}
			}
			if minNodeIndex != i {
				result++
				nextNode[minNodeIndex].Val, nextNode[i].Val = nextNode[i].Val, nextNode[minNodeIndex].Val
			}
		}
		nextNode, prevNode = prevNode, nextNode
		nextNode = nextNode[:0]
	}
	return result
}
