package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const modulo = 1e+9 + 7

func add(next [][]int, x, y, val int) int {
	if y < 0 || len(next) <= y {
		return val
	}
	if x < 0 || len(next[y]) <= x {
		return val
	}
	next[y][x] += val
	next[y][x] %= modulo
	return 0
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][]int, m)
	next := make([][]int, m)
	for y := 0; y < m; y++ {
		dp[y] = make([]int, n)
		next[y] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	result := 0
	for ; maxMove > 0; maxMove-- {
		for y, col := range dp {
			for x := range col {
				if dp[y][x] != 0 {
					result += add(next, x-1, y, dp[y][x])
					result += add(next, x+1, y, dp[y][x])
					result += add(next, x, y-1, dp[y][x])
					result += add(next, x, y+1, dp[y][x])
					result %= modulo
					dp[y][x] = 0
				}
			}
		}
		next, dp = dp, next
		// debug(dp)
	}
	return result
}

func debug(dp [][]int) {
	for _, col := range dp {
		fmt.Printf("dp: %v\n", col)
	}
	fmt.Println()
}

func TestPath(t *testing.T) {
	assert := require.New(t)
	assert.Equal(6, findPaths(2, 2, 2, 0, 0))
}
