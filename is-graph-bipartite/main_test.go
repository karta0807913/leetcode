package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	queue := make([]int, 0)
	for idx := range graph {
		if len(graph[idx]) != 0 {
			queue = append(queue, idx)
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				// fmt.Println(color, node)
				for _, next := range graph[node] {
					switch color[next] {
					case 0:
						color[next] = 1 ^ color[node]
						queue = append(queue, next)
					case color[node]:
						return false
					}
				}
			}
		}
	}
	return true
}

func TestBipartite(t *testing.T) {
	assert.False(t, isBipartite([][]int{
		{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
	}))
	assert.False(t, isBipartite([][]int{
		{1, 2}, {2, 0}, {0, 1},
	}))

	assert.False(t, isBipartite([][]int{
		{1}, {0}, {3, 4}, {2, 4}, {2, 3},
	}))
}
