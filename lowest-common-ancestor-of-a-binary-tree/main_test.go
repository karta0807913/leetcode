package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // slower
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	left := lowestCommonAncestor(root.Left, p, q)
// 	right := lowestCommonAncestor(root.Right, p, q)
// 	if left != nil && right != nil {
// 		return root
// 	}
// 	if root == p || root == q {
// 		return root
// 	}
// 	if left == nil {
// 		return right
// 	}
// 	return left
// }

// early terminate
func recursive(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	count := 0
	if root == p || root == q {
		count += 1
	}
	left, ok := recursive(root.Left, p, q)
	if ok {
		return left, ok
	}
	right, ok := recursive(root.Right, p, q)
	if ok {
		return right, ok
	}
	if left != nil {
		count += 1
	}
	if right != nil {
		count += 1
	}
	if count == 0 {
		return nil, false
	}
	return root, count == 2
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	root, _ = recursive(root, p, q)
	return root
}
