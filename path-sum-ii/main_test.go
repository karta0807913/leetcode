package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root *TreeNode, target int, path *[]int) (ans [][]int) {
	if root == nil {
		return nil
	}
	target -= root.Val
	*path = append((*path), root.Val)
	defer func() {
		*path = (*path)[:len(*path)-1]
	}()
	if target == 0 && root.Left == nil && root.Right == nil {
		return [][]int{append([]int{}, (*path)...)}
	}
	ans = append(ans, recursive(root.Left, target, path)...)
	ans = append(ans, recursive(root.Right, target, path)...)
	return
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	path := []int{}
	return recursive(root, targetSum, &path)
}
