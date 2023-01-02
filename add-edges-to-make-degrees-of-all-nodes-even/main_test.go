package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func isPossible(n int, edges [][]int) bool {
	graph := make(map[int]map[int]bool)
	count := make(map[int]int)
	insertGraph := func(from, to int) {
		if graph[from] == nil {
			graph[from] = make(map[int]bool)
		}
		graph[from][to] = true
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		insertGraph(from, to)
		insertGraph(to, from)
		count[to]++
		count[from]++
	}
	oddNodes := []int{}
	for node, count := range count {
		if count%2 == 1 {
			oddNodes = append(oddNodes, node)
		}
	}
	if len(oddNodes) == 2 {
		if !graph[oddNodes[0]][oddNodes[1]] {
			return true
		}
		for i := 1; i <= n; i++ {
			if !graph[oddNodes[0]][i] && !graph[oddNodes[1]][i] {
				return true
			}
		}
		return false
	}
	if len(oddNodes) == 4 {
		a, b, c, d := oddNodes[0], oddNodes[1], oddNodes[2], oddNodes[3]
		return (!graph[a][b] && !graph[c][d]) ||
			(!graph[a][c] && !graph[b][d]) ||
			(!graph[a][d] && !graph[b][c])
	}
	return len(oddNodes) == 0
}

func TestIsPossible(t *testing.T) {
	assert := require.New(t)
	assert.False(isPossible(4, [][]int{
		{1, 2}, {2, 3}, {2, 4}, {3, 4},
	}))
	assert.True(isPossible(5, [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}, {2, 5},
	}))
	assert.True(isPossible(4, [][]int{
		{1, 2}, {3, 4},
	}))
	assert.False(isPossible(4, [][]int{
		{1, 2}, {1, 3}, {1, 4},
	}))
	assert.True(isPossible(4, [][]int{
		{4, 1}, {3, 2}, {2, 4}, {1, 3},
	}))
}
