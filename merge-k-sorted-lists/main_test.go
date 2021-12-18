package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		val, next,
	}
}

func ToArray(node *ListNode) []int {
	result := make([]int, 0)
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}

func FromArray(array []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, val := range array {
		p.Next = &ListNode{
			val, nil,
		}
		p = p.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	pointer := head
	var list_end = 0
	end := len(lists)
	for list_end < end {
		if lists[list_end] == nil {
			end -= 1
			lists[list_end], lists[end] = lists[end], lists[list_end]
		} else {
			list_end += 1
		}
	}
	if list_end == 0 {
		return nil
	}
	for list_end != 1 {
		min_idx := 0
		for idx, list := range lists[1:list_end] {
			if list.Val < lists[min_idx].Val {
				min_idx = idx + 1
			}
		}
		pointer.Next = lists[min_idx]
		lists[min_idx] = lists[min_idx].Next
		pointer = pointer.Next

		if lists[min_idx] == nil {
			list_end -= 1
			lists[min_idx], lists[list_end] = lists[list_end], lists[min_idx]
		}
	}
	pointer.Next = lists[0]
	return head.Next
}

func TestMerge(t *testing.T) {
	input := []*ListNode{
		FromArray([]int{1, 4, 5}),
		FromArray([]int{1, 3, 4}),
		FromArray([]int{2, 6}),
	}

	assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, ToArray(mergeKLists(input)), nil)

	assert.Empty(t, mergeKLists([]*ListNode{nil, nil}), nil)
}
