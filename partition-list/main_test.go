package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	l := &ListNode{}
	r := &ListNode{}
	lp := l
	rp := r

	for head != nil {
		next := head.Next
		if head.Val < x {
			lp.Next = head
			lp = lp.Next
		} else {
			rp.Next = head
			rp = rp.Next
		}
		head = next
	}
	rp.Next = nil
	lp.Next = r.Next
	return l.Next
}
