package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Iterator []*TreeNode

func NewIterator(node *TreeNode) (iterator Iterator) {
	pointer := &node
	for *pointer != nil {
		iterator = append(iterator, *pointer)
		pointer = &(*pointer).Left
	}
	return
}

func (iterator *Iterator) Next() *TreeNode {
	if len(*iterator) != 0 {
		n := len(*iterator) - 1
		node := (*iterator)[n]
		*iterator = (*iterator)[:n]
		pointer := &node.Right
		for *pointer != nil {
			*iterator = append(*iterator, *pointer)
			pointer = &(*pointer).Left
		}
		return node
	}
	return nil
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	left, right := NewIterator(root1), NewIterator(root2)
	lNode, rNode := left.Next(), right.Next()
	if lNode == nil {
		lNode, rNode, left, right = rNode, lNode, right, left
	}
	for lNode != nil {
		if rNode != nil {
			if lNode.Val < rNode.Val {
				ans = append(ans, lNode.Val)
				lNode = left.Next()
			} else {
				ans = append(ans, rNode.Val)
				rNode = right.Next()
			}
		} else {
			ans = append(ans, lNode.Val)
			lNode = left.Next()
		}
		if lNode == nil {
			lNode, rNode, left, right = rNode, lNode, right, left
		}
	}
	return
}
