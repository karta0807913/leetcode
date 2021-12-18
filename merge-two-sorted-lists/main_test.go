package main

import (
	"fmt"
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

func PrintNode(node *ListNode) {
	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func ToArray(node *ListNode) []int {
	result := make([]int, 0)
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p_1 := l1
	p_2 := l2
	head := &ListNode{0, nil}
	result := head
	for p_1 != nil && p_2 != nil {
		if p_1.Val < p_2.Val {
			result.Next = p_1
			p_1 = p_1.Next
		} else {
			result.Next = p_2
			p_2 = p_2.Next
		}
		result = result.Next
	}
	if p_1 == nil {
		result.Next = p_2
	} else {
		result.Next = p_1
	}
	return head.Next
}

func TestMerge(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, ToArray(mergeTwoLists(NewNode(1, NewNode(4, nil)), NewNode(2, NewNode(3, nil)))))
}
