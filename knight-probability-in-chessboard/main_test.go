package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func add(next [][]float64, x, y int, val float64) float64 {
    if y < 0 || len(next) <= y {
        return val
    }
    if x < 0 || len(next[y]) <= x {
        return val
    }
    next[y][x] += val
    return 0
}
func knightProbability(n int, k int, row int, column int) float64 {
    dp := make([][]float64, n)
    next := make([][]float64, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]float64, n)
        next[i] = make([]float64, n)
    }

    dp[row][column] = 1
    for k := k; k > 0; k-- {
        for y, column := range dp {
            for x := range column {
                if dp[y][x] != 0 {
                    add(next, x+1, y-2, dp[y][x]/8)
                    add(next, x-1, y-2, dp[y][x]/8)
                    add(next, x-2, y-1, dp[y][x]/8)
                    add(next, x-2, y+1, dp[y][x]/8)
                    add(next, x-1, y+2, dp[y][x]/8)
                    add(next, x+1, y+2, dp[y][x]/8)
                    add(next, x+2, y+1, dp[y][x]/8)
                    add(next, x+2, y-1, dp[y][x]/8)
                    dp[y][x] = 0
                }
            }
        }
        next, dp = dp, next
    }
    // debug(dp)

    stay := float64(0)
    for _, column := range dp {
        for _, val := range column {
            stay += float64(val)
        }
    }
    return float64(stay)
}

func debug(dp [][]int) {
	for _, col := range dp {
		fmt.Printf("dp: %v\n", col)
	}
	fmt.Println()
}

func TestProbability(t *testing.T) {
	assert := require.New(t)
	assert.Equal(0.06250, knightProbability(3, 2, 0, 0))
}
