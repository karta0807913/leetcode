package main

import (
	"container/heap"
	"math"
	"sort"
)

type EndPoint struct {
	Height int
	X      int
}

type EndpointHeap []EndPoint

func (h EndpointHeap) Len() int           { return len(h) }
func (h EndpointHeap) Less(i, j int) bool { return h[i].Height > h[j].Height }
func (h EndpointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EndpointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(EndPoint))
}

func (h *EndpointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i][0] < buildings[j][0]
	})
	endHeap := EndpointHeap{}
	heap.Push(&endHeap, EndPoint{
		X:      math.MaxInt,
		Height: 0,
	})
	lastHeight := 0
	lastEnd := math.MaxInt
	ans := [][]int{}
	for _, building := range buildings {
		start := building[0]
		end := building[1]
		height := building[2]
		for lastEnd < start {
			for endHeap[0].X <= lastEnd {
				heap.Pop(&endHeap)
			}
			if lastHeight != endHeap[0].Height {
				ans = append(ans, []int{lastEnd, endHeap[0].Height})
			}
			lastHeight = endHeap[0].Height
			lastEnd = endHeap[0].X
		}
		if height > lastHeight {
			if len(ans) != 0 && start == ans[len(ans)-1][0] {
				ans[len(ans)-1][1] = height
			} else {
				ans = append(ans, []int{start, height})
			}
			lastHeight = height
			lastEnd = end
		}
		heap.Push(&endHeap, EndPoint{
			X:      end,
			Height: height,
		})
	}
	for len(endHeap) != 1 {
		for endHeap[0].X <= lastEnd {
			heap.Pop(&endHeap)
		}
		if lastHeight != endHeap[0].Height {
			ans = append(ans, []int{lastEnd, endHeap[0].Height})
		}
		lastEnd = endHeap[0].X
		lastHeight = endHeap[0].Height
	}
	return ans
}