package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func distance(bomb1, bomb2 []int) float64 {
	x := float64(bomb1[0] - bomb2[0])
	y := float64(bomb1[1] - bomb2[1])
	return math.Sqrt(x*x + y*y)
}

func dfs(bombs [][]int, currentIndex int) int {
	ans := 1
	currentBomb := bombs[currentIndex]
	if currentBomb == nil {
		return 0
	}
	bombs[currentIndex] = nil
	for idx, bomb := range bombs {
		if idx == currentIndex || bomb == nil {
			continue
		}
		if distance(currentBomb, bomb) <= float64(currentBomb[2]) {
			ans += dfs(bombs, idx)
		}
	}
	return ans
}

func maximumDetonation(bombs [][]int) int {
	maxVal := 0
	for idx := range bombs {
		bombs := append([][]int{}, bombs...)
		if val := dfs(bombs, idx); val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func TestBomb(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		maximumDetonation([][]int{
			{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4},
		}),
		5,
	)
	assert.Equal(
		maximumDetonation([][]int{
			{2, 1, 3}, {6, 1, 4},
		}),
		2,
	)
}
