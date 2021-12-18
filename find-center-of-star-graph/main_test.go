package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func findCenter(edges [][]int) int {
// 	counter := make([]int, len(edges)+1)
// 	for _, edge := range edges {
// 		counter[edge[0]-1] += 1
// 		counter[edge[1]-1] += 1
// 		if counter[edge[0]-1] == len(edges) {
// 			return edge[0]
// 		}
// 		if counter[edge[1]-1] == len(edges) {
// 			return edge[1]
// 		}
// 	}
// 	return -1
// }

func findCenter(edges [][]int) int {
	a := edges[0]
	b := edges[1]
	if a[0] == b[0] {
		return a[0]
	}
	if a[1] == b[0] {
		return a[1]
	}
	if a[0] == b[1] {
		return a[0]
	}
	if a[1] == b[1] {
		return a[1]
	}
	return -1
}

func TestCenter(t *testing.T) {
	assert.Equal(t, 3, findCenter([][]int{{1, 3}, {2, 3}}), nil)
	assert.Equal(t, 2, findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}), nil)
	assert.Equal(t, 1, findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}), nil)
}
