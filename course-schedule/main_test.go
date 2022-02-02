package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	incomes := make([]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[0]] = append(edges[pre[0]], pre[1])
		incomes[pre[1]] += 1
	}
	nodes := make([]int, 0, numCourses)
	for nodeIdx, income := range incomes {
		if income == 0 {
			nodes = append(nodes, nodeIdx)
			numCourses -= 1
		}
	}
	for len(nodes) != 0 && numCourses >= 0 {
		node := nodes[0]
		nodes = nodes[1:]
		for _, edge := range edges[node] {
			incomes[edge] -= 1
			if incomes[edge] <= 0 {
				nodes = append(nodes, edge)
				numCourses -= 1
			}
		}
	}
	// fmt.Printf("numCourses: %v\n", numCourses)
	return numCourses == 0
}

func TestCan(t *testing.T) {
	assert.True(t, canFinish(2, [][]int{
		{1, 0},
	}))
	assert.False(t, canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	}))
	assert.True(t, canFinish(3, [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}))
	assert.False(t, canFinish(3, [][]int{
		{0, 1},
		{0, 2},
		{2, 0},
	}))
}
