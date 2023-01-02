package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func search(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if root.Val == k {
		return true
	} else if root.Val < k {
		return search(root.Right, k)
	}
	return search(root.Left, k)
}

func iterate(current **TreeNode, root *TreeNode, k int) bool {
	if *current == nil {
		return false
	}
	tmpNode := *current
	tmpK := k - tmpNode.Val
	if tmpK == tmpNode.Val {
		return false
	}
	*current = nil
	if search(root, tmpK) {
		return true
	}
	if tmpK < tmpNode.Val {
		if search(tmpNode.Right, tmpK) {
			return true
		}
	} else {
		if search(tmpNode.Left, tmpK) {
			return true
		}
	}
	*current = tmpNode
	return iterate(&(*current).Left, root, k) || iterate(&(*current).Right, root, k)
}

func findTarget(root *TreeNode, k int) bool {
	return iterate(&root.Left, root, k) || iterate(&root.Right, root, k)
}
