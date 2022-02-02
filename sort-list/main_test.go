package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSort(head *ListNode, length int) *ListNode {
	if length == 1 {
		head.Next = nil
		return head
	}
	n := length / 2
	right := head
	for n := n; n > 0; n-- {
		right = right.Next
	}
	return merge(mergeSort(head, n), mergeSort(right, length-n))
}

func merge(left, right *ListNode) *ListNode {
	var head *ListNode
	runner := &head
	for {
		if left == nil {
			left = right
			break
		}
		if right == nil {
			break
		}
		if right.Val < left.Val {
			*runner = right
			runner = &(*runner).Next
			right = right.Next
		} else {
			*runner = left
			runner = &(*runner).Next
			left = left.Next
		}
	}
	*runner = left
	return head
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	pointer := head
	for pointer != nil {
		length += 1
	}
	return mergeSort(head, length)
}
