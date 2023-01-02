package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const mod = 1e+9 + 7

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	sortedEfficiencyAray := make([]int, 0, len(efficiency))
	for i := range efficiency {
		sortedEfficiencyAray = append(sortedEfficiencyAray, i)
	}
	sort.Slice(sortedEfficiencyAray, func(i, j int) bool {
		return efficiency[sortedEfficiencyAray[i]] >
			efficiency[sortedEfficiencyAray[j]]
	})
	totalSpeed := 0
	totalEfficiency := 0
	intHeap := make(IntHeap, 0, n)
	result := uint64(0)

	for len(sortedEfficiencyAray) != 0 {
		index := sortedEfficiencyAray[0]
		totalSpeed += speed[index]
		totalEfficiency = efficiency[index]
		heap.Push(&intHeap, speed[index])
		if len(intHeap) > k {
			poped := heap.Pop(&intHeap).(int)
			totalSpeed -= poped
		}
		tmp := (uint64(totalSpeed) * uint64(totalEfficiency))
		if tmp > result {
			result = tmp
		}
		sortedEfficiencyAray = sortedEfficiencyAray[1:]
	}
	return int(result % mod)
}
