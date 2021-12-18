package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectTwoNode(left, right *Node) {
	if left.Right == nil {
		return
	}
	if left.Right.Next != nil {
		return
	}
	left.Right.Next = right.Left
	connectTwoNode(left.Right, right.Left)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	connect(root.Left)
	connect(root.Right)
	connectTwoNode(root.Left, root.Right)
	return root
}
