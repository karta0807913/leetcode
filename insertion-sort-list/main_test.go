package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	current := &head
	for *current != nil {
		pointer := &head
		for *pointer != *current && (*pointer).Val < (*current).Val {
			pointer = &(*pointer).Next
		}
		if *pointer != *current {
			*current, (*current).Next, *pointer = (*current).Next, *pointer, *current
		} else {
			current = &(*current).Next
		}
	}
	return head
}
