package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func min(mat [][]int, x, y int) int {
	min := get(mat, x-1, y)
	if tmp := get(mat, x+1, y); min > tmp {
		min = tmp
	}
	if tmp := get(mat, x, y-1); min > tmp {
		min = tmp
	}
	if tmp := get(mat, x, y+1); min > tmp {
		min = tmp
	}
	return min
}

func get(mat [][]int, x, y int) int {
	if y < 0 || len(mat) <= y {
		return 99999999
	}
	if x < 0 || len(mat[y]) <= x {
		return 99999999
	}
	return mat[y][x]
}

func updateMatrix(mat [][]int) [][]int {
	searchPool := make([]int, 0, len(mat)*len(mat[0]))
	for y := range mat {
		for x := range mat[y] {
			c := y*len(mat[y]) + x
			if mat[y][x] != 0 {
				searchPool = append(searchPool, c)
			}
		}
	}
	fmt.Printf("searchPool: %v\n", searchPool)
	distance := 2
	for len(searchPool) != 0 {
		left := 0
		for right := 0; right < len(searchPool); right++ {
			y, x := searchPool[right]/len(mat[0]), searchPool[right]%len(mat[0])
			next := min(mat, x, y) + 1
			if next >= distance {
				mat[y][x] = next
				searchPool[left], searchPool[right] =
					searchPool[right], searchPool[left]
				left++
			}
		}
		searchPool = searchPool[:left]
		fmt.Printf("mat: %v\n", mat)
		distance++
	}
	return mat
}

func TestMatrix(t *testing.T) {
	assert := require.New(t)
	assert.Equal([][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}, updateMatrix([][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}))
	assert.Equal([][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 2, 1},
	}, updateMatrix([][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
	}))
	assert.Equal([][]int{
		{1, 2, 1}, {0, 1, 0}, {0, 0, 0},
	}, updateMatrix([][]int{
		{1, 1, 1}, {0, 1, 0}, {0, 0, 0},
	}))
	assert.Equal([][]int{
		{1, 2, 1}, {0, 1, 0}, {1, 2, 1},
	}, updateMatrix([][]int{
		{1, 1, 1}, {0, 1, 0}, {1, 1, 1},
	}))
	assert.Equal([][]int{
		{4, 3, 2, 3, 4},
		{3, 2, 1, 2, 3},
		{2, 1, 0, 1, 2},
		{3, 2, 1, 2, 3},
		{4, 3, 2, 3, 4},
	}, updateMatrix([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
}
