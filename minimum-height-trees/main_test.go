package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	tree := make([][]int, n)
	incomes := make([]int, n)
	for _, edge := range edges {
		left, right := edge[0], edge[1]
		tree[left] = append(tree[left], right)
		tree[right] = append(tree[right], left)
		incomes[left] += 1
		incomes[right] += 1
	}

	runner := []int{}
	next := []int{}
	for idx, references := range incomes {
		if references == 1 {
			runner = append(runner, idx)
			incomes[idx] -= 1
		}
	}
	for {
		for _, node := range runner {
			for _, edge := range tree[node] {
				incomes[edge] -= 1
				if incomes[edge] == 1 {
					next = append(next, edge)
				}
			}
		}
		if len(next) == 0 {
			break
		}
		runner, next = next, []int{}
	}
	return runner
}

func TestTree(t *testing.T) {
	assert.Equal(t, []int{
		1,
	}, findMinHeightTrees(4, [][]int{
		{1, 0}, {1, 2}, {1, 3},
	}))
}
