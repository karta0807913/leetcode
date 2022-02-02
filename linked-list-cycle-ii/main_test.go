package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				slow, fast = slow.Next, fast.Next
			}
			return fast
		}
	}
}
