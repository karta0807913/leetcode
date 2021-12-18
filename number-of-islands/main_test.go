package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Spot struct {
	x int
	y int
}

func expand(x, y int, landIdx byte, grid [][]byte) {
	searchingSpot := []Spot{{
		x: x,
		y: y,
	}}
	for ; len(searchingSpot) != 0; searchingSpot = searchingSpot[1:] {
		val := searchingSpot[0]
		if val.y < 0 || val.y == len(grid) || val.x < 0 || val.x == len(grid[0]) {
			continue
		}
		if grid[val.y][val.x] == '1' {
			grid[val.y][val.x] = landIdx
			searchingSpot = append(searchingSpot, Spot{x: val.x + 1, y: val.y})
			searchingSpot = append(searchingSpot, Spot{x: val.x - 1, y: val.y})
			searchingSpot = append(searchingSpot, Spot{x: val.x, y: val.y + 1})
			searchingSpot = append(searchingSpot, Spot{x: val.x, y: val.y - 1})
		}
	}
}

func printGrid(grid [][]byte) {
	for _, t := range grid {
		for _, val := range t {
			fmt.Print(string(val) + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func numIslands(grid [][]byte) int {
	landIdx := byte('1' + 1)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '1' {
				expand(x, y, landIdx, grid)
				// printGrid(grid)
				landIdx += 1
			}
		}
	}
	return int(landIdx) - ('1' + 1)
}

func TestNum(t *testing.T) {
	assert.Equal(t, 1, numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}), nil)

	assert.Equal(t, 3, numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}), nil)
}
