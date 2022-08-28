package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root *TreeNode, voyage *[]int) []int {
	if root == nil {
		return nil
	}
	if (*voyage)[0] != root.Val {
		return nil
	}
	(*voyage) = (*voyage)[1:]
	if root.Left != nil && root.Left.Val ==  {
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	result := recursive(root, &voyage)
	if len(voyage) != 0 {
		return []int{-1}
	}
	return result
}
