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

func ToArray(node *ListNode) []int {
	result := make([]int, 0)
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	newList := &ListNode{-1, nil}
	foot := head
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = newList.Next
		newList.Next = tmp
	}
	foot.Next = nil
	return newList.Next, foot
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	newList := &ListNode{-1, head}
	idx := 1
	top := newList
	for head != nil {
		head = head.Next
		idx += 1
		if idx == k && head != nil {
			idx = 1
			tmp := head
			head = head.Next
			tmp.Next = nil
			start, end := reverse(top.Next)
			top.Next = start
			top = end
			top.Next = head
		}
	}
	return newList.Next
}

func TestReverse(t *testing.T) {
	input := NewNode(1, NewNode(2, NewNode(3, NewNode(4, nil))))
	head, foot := reverse(input)
	assert.Equal(t, []int{4, 3, 2, 1}, ToArray(head), nil)
	assert.Equal(t, 1, foot.Val, nil)

	input = NewNode(1, NewNode(2, NewNode(3, NewNode(4, NewNode(5, nil)))))
	result := reverseKGroup(input, 3)
	assert.Equal(t, []int{3, 2, 1, 4, 5}, ToArray(result), nil)

	input = NewNode(1, NewNode(2, NewNode(3, NewNode(4, nil))))
	result = reverseKGroup(input, 2)
	assert.Equal(t, []int{2, 1, 4, 3}, ToArray(result), nil)

	input = NewNode(1, NewNode(2, NewNode(3, NewNode(4, nil))))
	result = reverseKGroup(input, 3)
	assert.Equal(t, []int{3, 2, 1, 4}, ToArray(result), nil)
}
