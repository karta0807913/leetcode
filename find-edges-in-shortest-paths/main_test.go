package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

type Edge struct {
	Current int
	Cost    int
	Next    int
}

type Step struct {
	Edge
}

// An EdgeHeap is a min-heap of ints.
type EdgeHeap []Step

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x any) {
	*h = append(*h, x.(Step))
}

func (h *EdgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Node struct {
	Edge
	Parent []int
}

func findAnswer(n int, edges [][]int) []bool {
	costMap := make([]Node, n)
	graph := make([][]Edge, n)
	indexMap := make(map[[2]int]int)
	for idx, edge := range edges {
		indexMap[[2]int{edge[0], edge[1]}] = idx
		indexMap[[2]int{edge[1], edge[0]}] = idx
		graph[edge[0]] = append(graph[edge[0]], Edge{
			Current: edge[0],
			Next:    edge[1],
			Cost:    edge[2],
		})
		graph[edge[1]] = append(graph[edge[1]], Edge{
			Current: edge[1],
			Next:    edge[0],
			Cost:    edge[2],
		})
	}
	eHeap := EdgeHeap{
		Step{
			Edge: Edge{
				Current: 0,
				Next:    0,
				Cost:    0,
			},
		},
	}

	for len(eHeap) != 0 {
		step := heap.Pop(&eHeap).(Step)
		if cost := costMap[step.Next]; cost.Cost != 0 {
			if step.Cost == cost.Cost {
				cost.Parent = append(cost.Parent, step.Current)
				costMap[step.Next] = cost
			}
			continue
		}
		costMap[step.Next].Cost = step.Cost
		costMap[step.Next].Parent = append(costMap[step.Next].Parent, step.Current)
		if step.Next == n-1 {
			continue
		}
		for _, edge := range graph[step.Next] {
			heap.Push(&eHeap, Step{
				Edge: Edge{
					Current: step.Next,
					Next:    edge.Next,
					Cost:    edge.Cost + step.Cost,
				},
			})
		}
	}
	ans := make([]bool, len(edges))
	var recursive func(node Node)
	costMap[0].Parent = nil
	recursive = func(node Node) {
		for _, parent := range node.Parent {
			idx := [2]int{node.Current, parent}
			if ans[indexMap[idx]] {
				continue
			}
			ans[indexMap[idx]] = true
			recursive(costMap[parent])
		}
	}
	for i := range costMap {
		costMap[i].Current = i
	}
	recursive(costMap[n-1])
	return ans
}

func TestFindAnswer(t *testing.T) {
	assert := require.New(t)

	assert.Equal(
		[]bool{
			false, false, false, false, false, false, true, false, false, false,
		},
		findAnswer(
			7,
			[][]int{
				{2, 4, 4},
				{5, 4, 9},
				{0, 2, 6},
				{6, 2, 1},
				{3, 6, 3},
				{1, 3, 6},
				{6, 0, 4},
				{0, 4, 5},
				{1, 0, 1},
				{3, 5, 2},
			},
		),
	)
	assert.Equal(
		[]bool{
			true, true, true, false, true, true, true, false,
		},
		findAnswer(
			6,
			[][]int{
				{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2},
			},
		),
	)
}
