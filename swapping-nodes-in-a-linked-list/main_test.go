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

func swapNodes(head *ListNode, k int) *ListNode {
	newList := &ListNode{-1, head}
	// a_before := newList
	a_node := newList.Next
	for idx := 1; idx != k; idx++ {
		// a_before = a_node
		a_node = a_node.Next
	}
	ancher := a_node
	// b_before := newList
	b_node := newList
	for ancher != nil {
		ancher = ancher.Next
		// b_before = b_node
		b_node = b_node.Next
	}
	a_node.Val, b_node.Val = b_node.Val, a_node.Val
	return newList.Next
}

func TestSwap(t *testing.T) {
	input := FromArray([]int{1, 2, 3, 4, 5})
	result := swapNodes(input, 5)
	assert.Equal(t, []int{5, 2, 3, 4, 1}, ToArray(result), nil)

	input = FromArray([]int{1})
	result = swapNodes(input, 1)
	assert.Equal(t, []int{1}, ToArray(result), nil)

	input = FromArray([]int{1, 2, 3})
	result = swapNodes(input, 2)
	assert.Equal(t, []int{1, 2, 3}, ToArray(result), nil)

	input = FromArray([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	result = swapNodes(input, 5)
	fmt.Println(ToArray(result))
	assert.Equal(t, []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}, ToArray(result), nil)
}
