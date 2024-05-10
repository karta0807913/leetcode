package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	nodeCount := make([]int, n)
	nodeDistance := make([]int, n)
	ans := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	{
		var recursive func(nodeIdx int, parent int) (distance int, children int)
		recursive = func(nodeIdx, parent int) (distance int, children int) {
			for _, child := range graph[nodeIdx] {
				if child == parent {
					continue
				}
				_d, _c := recursive(child, nodeIdx)
				// 1 is child itself
				distance += _c + _d + 1
				children += _c + 1
			}
			nodeCount[nodeIdx] = children
			nodeDistance[nodeIdx] = distance
			return distance, children
		}
		recursive(0, -1)
	}

	var recursive func(nodeIdx int, parent int, nodeCount int, nodeDistance int)
	recursive = func(nodeIdx, parent, otherNum, otherDistances int) {
		distance := otherNum + otherDistances
		_otherDistances := distance
		for _, child := range graph[nodeIdx] {
			if child == parent {
				continue
			}
			distance += nodeDistance[child] + nodeCount[child] + 1
			_otherDistances += nodeDistance[child] + nodeCount[child] + 1
		}
		ans[nodeIdx] = distance
		for _, child := range graph[nodeIdx] {
			if child == parent {
				continue
			}
			recursive(
				child,
				nodeIdx,
				n-(nodeCount[child]+1),
				_otherDistances-(nodeDistance[child]+nodeCount[child]+1),
			)
		}
	}
	recursive(0, -1, 0, 0)

	return ans
}

func TestSumOfDistancesInTree(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{
			8, 12, 6, 10, 10, 10,
		},
		sumOfDistancesInTree(
			6,
			[][]int{
				{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5},
			},
		),
	)
}
