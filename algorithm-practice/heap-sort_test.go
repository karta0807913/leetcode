package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := IntHeap{2, 3}
	Push(&h, 1)
	assert.Equal(t, []int{1, 3, 2}, []int(h))
	assert.Equal(t, 1, Pop(&h))
	assert.Equal(t, []int{2, 3}, []int(h))
	question := []int{1, 4, 2, 456, 1, 2}
	h = append([]int{}, question...)
	Init(&h)
	answer := []int{}
	for h.Len() != 0 {
		// fmt.Printf("h: %v\n", h)
		answer = append(answer, Pop(&h).(int))
	}
	sort.Ints(question)
	assert.Equal(t, question, answer)

	for testsCount := 0; testsCount < 10; testsCount++ {
		question = []int{}
		for idx := 0; idx < 5000; idx++ {
			question = append(question, rand.Int())
		}
		h = append([]int{}, question[:len(question)-20]...)
		Init(&h)
		for idx := len(question) - 20; idx < len(question); idx++ {
			Push(&h, question[idx])
		}
		answer = []int{}
		for h.Len() != 0 {
			answer = append(answer, Pop(&h).(int))
		}
		sort.Ints(question)
		assert.Equal(t, question, answer)
	}
}
