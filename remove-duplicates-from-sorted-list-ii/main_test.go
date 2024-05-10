package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pointer := &head
	deleteCurrent := false
	for (*pointer).Next != nil {
		if (*pointer).Next.Val == (*pointer).Val {
			(*pointer).Next = (*pointer).Next.Next
			deleteCurrent = true
		} else if deleteCurrent {
			deleteCurrent = false
			*pointer = (*pointer).Next
		} else {
			pointer = &(*pointer).Next
		}
	}
	if deleteCurrent {
		*pointer = (*pointer).Next
	}
	return head
}

func arr2list(arr []int) *ListNode {
	head := &ListNode{}
	pointer := head
	for _, val := range arr {
		pointer.Next = &ListNode{Val: val}
		pointer = pointer.Next
	}
	return head.Next
}

func list2arr(list *ListNode) []int {
	ans := []int{}
	for list != nil {
		ans = append(ans, list.Val)
		list = list.Next
	}
	return ans
}

func TestDelete(t *testing.T) {
	assert.Equal(t, []int{
		1, 2, 5,
	}, list2arr(deleteDuplicates(arr2list([]int{
		1, 2, 3, 3, 4, 4, 5,
	}))))
	assert.Equal(t, []int{
		2, 3,
	}, list2arr(deleteDuplicates(arr2list([]int{
		1, 1, 1, 2, 3,
	}))))
	assert.Equal(t, []int{}, list2arr(deleteDuplicates(arr2list([]int{
		1, 1, 1, 1, 1,
	}))))
}
