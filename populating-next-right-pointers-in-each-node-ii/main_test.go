package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Due to the result, the double pointer version is faster.
 **/

// double pointer version
func connect(root *Node) *Node {
	current := root
	var next *Node
	for current != nil {
		pointer := &next
		for current != nil {
			if current.Left != nil {
				*pointer = current.Left
				pointer = &(*pointer).Next
			}
			if current.Right != nil {
				*pointer = current.Right
				pointer = &(*pointer).Next
			}
			current = current.Next
		}
		current = next
		next = nil
	}
	return root
}

// // normal dummy head version
// func connect(root *Node) *Node {
// 	current := root
// 	next := &Node{}
// 	for current != nil {
// 		pointer := next
// 		for current != nil {
// 			if current.Left != nil {
// 				pointer.Next = current.Left
// 				pointer = pointer.Next
// 			}
// 			if current.Right != nil {
// 				pointer.Next = current.Right
// 				pointer = pointer.Next
// 			}
// 			current = current.Next
// 		}
// 		current = next.Next
// 		next = &Node{}
// 	}
// 	return root
// }
