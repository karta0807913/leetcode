package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	edges := make([][]int, numCourses)
	incomes := make([]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		incomes[pre[0]] += 1
	}
	queue := make([]int, 0, numCourses)
	for idx, refer := range incomes {
		if refer == 0 {
			queue = append(queue, idx)
			ans = append(ans, idx)
		}
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, edge := range edges[node] {
			incomes[edge] -= 1
			if incomes[edge] == 0 {
				queue = append(queue, edge)
				ans = append(ans, edge)
			} else if incomes[edge] < 0 {
				return []int{}
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}

func TestOrder(t *testing.T) {
	assert.Equal(t, []int{
		0, 1,
	}, findOrder(2, [][]int{{1, 0}}))
	assert.Equal(t, []int{}, findOrder(3, [][]int{
		{1, 0},
		{0, 1},
		{1, 2},
	}))
}
