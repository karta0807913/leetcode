package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	fast := head
	slow := &head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = &(*slow).Next
		fast = fast.Next
	}
	*slow = (*slow).Next
	return head
}
