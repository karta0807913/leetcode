package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

type Pair struct {
	Val  int
	Left bool
}

type IntHeap []Pair

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	if h[i].Val < h[j].Val {
		return true
	}
	if h[i].Val == h[j].Val && h[i].Left {
		return true
	}
	return false
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	result := int64(0)
	var intHeap IntHeap
	var remains []int
	if candidates*2 < len(costs) {
		intHeap = make(IntHeap, 0, candidates*2)
		remains = costs[candidates : len(costs)-candidates]
		for i := 0; i < candidates; i++ {
			intHeap = append(intHeap, Pair{Left: true, Val: costs[i]})
		}
		for i := len(costs) - candidates; i < len(costs); i++ {
			intHeap = append(intHeap, Pair{Left: false, Val: costs[i]})
		}
	} else {
		intHeap = make(IntHeap, 0, len(costs))
		for _, cost := range costs {
			intHeap = append(intHeap, Pair{Val: cost})
		}
	}
	heap.Init(&intHeap)
	for k != 0 {
		cost := heap.Pop(&intHeap).(Pair)
		if len(remains) != 0 {
			if cost.Left {
				heap.Push(&intHeap, Pair{Val: remains[0], Left: true})
				remains = remains[1:]
			} else {
				heap.Push(&intHeap, Pair{Val: remains[len(remains)-1], Left: false})
				remains = remains[:len(remains)-1]
			}
		}
		result += int64(cost.Val)
		k--
	}
	return result
}

func TestCost(t *testing.T) {
	assert := require.New(t)
	assert.Equal(int64(11), totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	assert.Equal(int64(4), totalCost([]int{1, 2, 4, 1}, 3, 3))
}
