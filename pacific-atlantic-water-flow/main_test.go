package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const PACIFIC_BIT = 1
const ATLANTIC_BIT = 2

func pacificAtlanticRecursive(heights, flowMap [][]int, colIdx, rowIdx int, state int) {
	currentHeight := heights[colIdx][rowIdx]
	if flowMap[colIdx][rowIdx]&state != 0 {
		return
	}
	flowMap[colIdx][rowIdx] |= state

	fmt.Printf("flowMap: %v\n", flowMap)
	if colIdx != 0 && currentHeight <= heights[colIdx-1][rowIdx] {
		pacificAtlanticRecursive(heights, flowMap, colIdx-1, rowIdx, state)
	}
	if rowIdx != 0 && currentHeight <= heights[colIdx][rowIdx-1] {
		pacificAtlanticRecursive(heights, flowMap, colIdx, rowIdx-1, state)
	}
	if colIdx != len(heights)-1 && currentHeight <= heights[colIdx+1][rowIdx] {
		pacificAtlanticRecursive(heights, flowMap, colIdx+1, rowIdx, state)
	}
	if rowIdx != len(heights[0])-1 && currentHeight <= heights[colIdx][rowIdx+1] {
		pacificAtlanticRecursive(heights, flowMap, colIdx, rowIdx+1, state)
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	flowMap := make([][]int, len(heights))
	for colIdx, _ := range heights {
		flowMap[colIdx] = make([]int, len(heights[colIdx]))
	}
	for colIdx := range heights {
		pacificAtlanticRecursive(heights, flowMap, colIdx, 0, PACIFIC_BIT)
		pacificAtlanticRecursive(heights, flowMap, colIdx, len(heights[0])-1, ATLANTIC_BIT)
	}
	for rowIdx := range heights[0] {
		pacificAtlanticRecursive(heights, flowMap, 0, rowIdx, PACIFIC_BIT)
		pacificAtlanticRecursive(heights, flowMap, len(heights)-1, rowIdx, ATLANTIC_BIT)
	}
	fmt.Printf("flowMap: %v\n", flowMap)
	answer := make([][]int, 0)
	for colIdx, column := range heights {
		for rowIdx := range column {
			if flowMap[colIdx][rowIdx]&3 == (PACIFIC_BIT | ATLANTIC_BIT) {
				answer = append(answer, []int{colIdx, rowIdx})
			}
		}
	}
	return answer
}

func TestWater(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}, pacificAtlantic([][]int{
		{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4},
	}))
	assert.ElementsMatch(t, [][]int{
		{0, 0}, {0, 1}, {1, 0}, {1, 1},
	}, pacificAtlantic([][]int{
		{2, 1}, {1, 2},
	}))
	assert.ElementsMatch(t, [][]int{
		{0, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 0}, {2, 1}, {2, 2}, {2, 3}, {3, 0}, {3, 1}, {3, 2}, {3, 3},
	}, pacificAtlantic([][]int{
		{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7},
	}))

	assert.ElementsMatch(t, [][]int{
		{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1},
	}, pacificAtlantic([][]int{
		{1, 1}, {1, 1}, {1, 1},
	}))

	assert.ElementsMatch(t, [][]int{
		{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2},
	}, pacificAtlantic([][]int{
		{10, 10, 10}, {10, 1, 10}, {10, 10, 10},
	}))

	assert.ElementsMatch(t, [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {1, 0}, {1, 2}, {1, 3}, {1, 5}, {2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5},
	}, pacificAtlantic([][]int{
		{3, 3, 3, 3, 3, 3}, {3, 0, 3, 3, 0, 3}, {3, 3, 3, 3, 3, 3},
	}))
}
