package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Searched bool
	Depth    int
	Next     []int
}

func iterate(graph []Node, startIdx int) bool {
	nextList := []int{startIdx}
	graph[startIdx].Depth = 1
	for idx := 0; idx < len(nextList); idx++ {
		node := &graph[nextList[idx]]
		for _, val := range node.Next {
			if graph[val].Searched {
				continue
			}
			if graph[val].Depth != 0 {
				if graph[val].Depth&1 == node.Depth&1 {
					return false
				}
			} else {
				graph[val].Depth = node.Depth + 1
				nextList = append(nextList, val)
			}
		}
	}
	for _, idx := range nextList {
		graph[idx].Searched = true
	}
	return true
}

// Bipartite Graph Recognition Problem
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([]Node, n+1)
	for _, dislike := range dislikes {
		graph[dislike[0]].Next = append(graph[dislike[0]].Next, dislike[1])
	}
	for idx := range graph {
		if !graph[idx].Searched {
			if !iterate(graph, idx) {
				return false
			}
			graph[idx].Searched = true
		}
	}
	return true
}

func TestPossible(t *testing.T) {
	assert.True(t, possibleBipartition(4, [][]int{
		{1, 2}, {1, 3}, {2, 4},
	}))
	assert.False(t, possibleBipartition(3, [][]int{
		{1, 2}, {1, 3}, {2, 3},
	}))
	assert.False(t, possibleBipartition(5, [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5},
	}))
	assert.True(t, possibleBipartition(10, [][]int{
		{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8},
	}))
	assert.True(t, possibleBipartition(4, [][]int{
		{3, 4}, {1, 2}, {2, 4},
	}))
}
