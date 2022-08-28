package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	currentPosition := startFuel
	intHeap := &IntHeap{}
	result := 0
	stations = append(stations, []int{target, 0})
	for i := 0; i < len(stations); i++ {
		station := stations[i]
		if station[0] <= currentPosition {
			heap.Push(intHeap, station[1])
		} else {
			i--
			// fmt.Printf("station[0]: %v\n", station[0])
			if intHeap.Len() == 0 {
				return -1
			}
			currentPosition += heap.Pop(intHeap).(int)
			// fmt.Printf("currentPosition: %v\n", currentPosition)
			result += 1
		}
	}
	return result
}

func TestStops(t *testing.T) {
	assert := require.New(t)
	assert.Equal(4, minRefuelStops(1000, 299, [][]int{
		{13, 21},
		{26, 115},
		{100, 47},
		{225, 99},
		{299, 141},
		{444, 198},
		{608, 190},
		{636, 157},
		{647, 255},
		{841, 123},
	}))
	assert.Equal(3, minRefuelStops(100, 25, [][]int{
		{25, 25},
		{50, 25},
		{75, 25},
	}))
	assert.Equal(2, minRefuelStops(100, 10, [][]int{
		{10, 60},
		{20, 30},
		{30, 30},
		{60, 40},
	}))
	assert.Equal(0, minRefuelStops(100, 100, [][]int{
		{10, 60},
		{20, 30},
		{30, 30},
		{60, 40},
	}))
}
