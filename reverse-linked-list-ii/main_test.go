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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	newList := &ListNode{-1, head}
	left_pointer := newList
	right_pointer := newList
	for idx := 1; idx < left; idx++ {
		left_pointer = left_pointer.Next
		right_pointer = right_pointer.Next
	}

	for idx := left - 1; idx < right; idx++ {
		right_pointer = right_pointer.Next
	}
	next := right_pointer.Next
	right_pointer.Next = nil
	l, r := reverse(left_pointer.Next)
	r.Next = next
	left_pointer.Next = l
	return newList.Next
}

func TestReverse(t *testing.T) {
	input := NewNode(1, NewNode(2, NewNode(3, NewNode(4, NewNode(5, nil)))))
	result := reverseBetween(input, 2, 4)
	assert.Equal(t, []int{1, 4, 3, 2, 5}, ToArray(result), nil)
}
