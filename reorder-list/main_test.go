package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	for ; list != nil; list = list.Next {
		fmt.Printf("%d, ", list.Val)
	}
	fmt.Println()
}

func reorderList(head *ListNode) {
	n := 0
	var pointer *ListNode
	for pointer = head; pointer != nil; pointer = pointer.Next {
		n++
	}
	secondHead := new(*ListNode)
	n -= n / 2
	pointer = head
	for ; n > 1; n-- {
		pointer = pointer.Next
	}
	pointer, pointer.Next = pointer.Next, nil
	for pointer != nil {
		*secondHead, pointer.Next, pointer = pointer, *secondHead, pointer.Next
	}
	// printList(secondHead)
	pointer = *secondHead
	for pointer != nil {
		head.Next, pointer.Next, head, pointer = pointer, head.Next, head.Next, pointer.Next
	}
}
