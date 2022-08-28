package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func get(matrix [][]int, x, y int) (int, bool) {
	if x < 0 || x == len(matrix[0]) || y < 0 || y == len(matrix) {
		return -1, false
	}
	return matrix[y][x], true
}

func set(dp [][]int, x, y, z int) {
	if x < 0 || x == len(dp[0]) || y < 0 || y == len(dp) {
		return
	}
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	next := [][2]int{}
	for y, column := range matrix {
		dp[y] = make([]int, len(matrix[y]))
		for x := range column {
			if m, b := get(matrix, x+1, y); m > matrix[y][x] && b {
				continue
			}
			if m, b := get(matrix, x-1, y); m > matrix[y][x] && b {
				continue
			}
			if m, b := get(matrix, x, y+1); m > matrix[y][x] && b {
				continue
			}
			if m, b := get(matrix, x, y-1); m > matrix[y][x] && b {
				continue
			}
			next = append(next, [2]int{y, x})
			dp[y][x] = 1
		}
	}
	check := func(x, y, x1, y1 int) {
		if m, b := get(matrix, x1, y1); m < matrix[y][x] && b {
			if dp[y][x] >= dp[y1][x1] {
				dp[y1][x1] = dp[y][x] + 1
				next = append(next, [2]int{y1, x1})
			}
		}
	}
	m := 0
	for len(next) != 0 {
		x := next[0][1]
		y := next[0][0]
		// fmt.Printf("next: %v\n", next)
		next = next[1:]

		m = max(dp[y][x], m)
		check(x, y, x+1, y)
		check(x, y, x-1, y)
		check(x, y, x, y+1)
		check(x, y, x, y-1)
		// for _, column := range dp {
		// 	fmt.Printf("dp[y]: %v\n", column)
		// }
		// fmt.Println()
	}

	return m
}

func TestPath(t *testing.T) {
	assert := require.New(t)
	// assert.Equal(4, longestIncreasingPath([][]int{
	// 	{3, 4, 5},
	// 	{3, 2, 6},
	// }))
	// assert.Equal(9, longestIncreasingPath([][]int{
	// 	{9, 2, 3},
	// 	{8, 1, 4},
	// 	{7, 6, 5},
	// }))
	// assert.Equal(9, longestIncreasingPath([][]int{
	// 	{1, 2, 3},
	// 	{6, 5, 4},
	// 	{7, 8, 9},
	// }))
	assert.Equal(6, longestIncreasingPath([][]int{
		{1, 6, 12, 1, 3},
		{8, 4, 6, 10, 5},
		{12, 11, 7, 12, 2},
		{2, 3, 4, 1, 13},
		{14, 6, 0, 14, 13},
	}))
	// assert.Equal(70, longestIncreasingPath([][]int{
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
	// 	{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
	// 	{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
	// 	{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
	// 	{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
	// 	{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }))

	// assert.Equal(140, longestIncreasingPath([][]int{
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
	// 	{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
	// 	{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
	// 	{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
	// 	{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
	// 	{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
	// 	{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
	// 	{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
	// 	{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
	// 	{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
	// 	{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
	// 	{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
	// 	{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }))
}
