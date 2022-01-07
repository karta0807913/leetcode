package main

import (
	"sort"
)

type HeapInterface interface {
	sort.Interface
	Pop() interface{}
	Push(val interface{})
}

func Init(h HeapInterface) {
	n := h.Len()
	// make the smallest value to the array[0]
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Pop(h HeapInterface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Push(h HeapInterface, key interface{}) {
	h.Push(key)
	n := h.Len()
	moveUP(h, n-1)
}

func moveUP(h HeapInterface, j int) {
	i := (j - 1) / 2 // parent
	// move up
	for j != i && h.Less(j, i) {
		h.Swap(i, j)
		j = i
		i = (j - 1) / 2 // parent
	}
}

func down(h HeapInterface, i0, n int) bool {
	i := i0

	for {
		// left node
		j := i*2 + 1
		if j >= n || j < 0 {
			break
		}

		// find the smallest value in left node and right node
		// j2 is the right node
		if j2 := j + 1; 0 <= j2 && j2 < n && h.Less(j2, j) {
			j = j2
		}

		// if the parent is the smallest, return
		if h.Less(i, j) {
			break
		}

		h.Swap(i, j)
		// goto next level
		i = j
	}

	// check subtree has been modified
	return i > i0
}
