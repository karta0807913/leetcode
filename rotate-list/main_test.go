package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	pointer := head
	tail := head
	length := 0
	for pointer != nil {
		length += 1
		tail = pointer
		pointer = pointer.Next
	}

	k %= length
	if k == 0 {
		return head
	}
	newHead := head
	for forward := length - k - 1; forward > 0; forward-- {
		newHead = newHead.Next
	}
	newHead.Next, newHead = nil, newHead.Next
	tail.Next = head

	return newHead
}

func TestRotate(t *testing.T) {
	List2Array := func(list *ListNode) []int {
		res := []int{}
		for list != nil {
			res = append(res, list.Val)
			list = list.Next
		}
		return res
	}
	Array2List := func(arr []int) *ListNode {
		head := &ListNode{}
		pointer := head
		for _, val := range arr {
			pointer.Next = &ListNode{
				Val: val,
			}
			pointer = pointer.Next
		}
		return head.Next
	}
	assert.Equal(t, []int{2, 1}, List2Array(rotateRight(Array2List([]int{1, 2}), 1)))
	assert.Equal(t, []int{1}, List2Array(rotateRight(Array2List([]int{1}), 1)))
	assert.Equal(t, []int{1, 2}, List2Array(rotateRight(Array2List([]int{1, 2}), 2)))
}
