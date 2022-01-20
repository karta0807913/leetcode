package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toBST(head **ListNode, start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (end + start) / 2
	left := toBST(head, start, mid-1)
	if *head == nil {
		return left
	}
	val := (*head).Val
	(*head) = (*head).Next
	return &TreeNode{
		Left:  left,
		Val:   val,
		Right: toBST(head, mid+1, end),
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	pointer := head
	n := 0
	for pointer != nil {
		n += 1
		pointer = pointer.Next
	}
	return toBST(&head, 0, n-1)
}
