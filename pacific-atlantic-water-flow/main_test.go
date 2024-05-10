package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const PACIFIC_BIT = 1
const ATLANTIC_BIT = 2

func recursive(heights, statusMap [][]int, x, y int, status int) {
	if statusMap[y][x]&status != 0 {
		return
	}
	statusMap[y][x] |= status

	if x > 0 && heights[y][x] <= heights[y][x-1] {
		recursive(heights, statusMap, x-1, y, status)
	}

	if y > 0 && heights[y][x] <= heights[y-1][x] {
		recursive(heights, statusMap, x, y-1, status)
	}

	if x+1 < len(heights[y]) && heights[y][x] <= heights[y][x+1] {
		recursive(heights, statusMap, x+1, y, status)
	}

	if y+1 < len(heights) && heights[y][x] <= heights[y+1][x] {
		recursive(heights, statusMap, x, y+1, status)
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	statusMap := make([][]int, 0, len(heights))
	for _, column := range heights {
		statusMap = append(statusMap, make([]int, len(column)))
	}
	for y := range heights {
		recursive(heights, statusMap, 0, y, PACIFIC_BIT)
		recursive(heights, statusMap, len(heights[0])-1, y, ATLANTIC_BIT)
	}
	for x := range heights[0] {
		recursive(heights, statusMap, x, 0, PACIFIC_BIT)
		recursive(heights, statusMap, x, len(heights)-1, ATLANTIC_BIT)
	}
	answer := [][]int{}
	for y, column := range heights {
		for x := range column {
			if statusMap[y][x] == (PACIFIC_BIT | ATLANTIC_BIT) {
				answer = append(answer, []int{y, x})
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
