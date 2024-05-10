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

func dp(cache [][]int, piles []int, M int, currentIndex int) int {
	if currentIndex >= len(piles) {
		return 0
	}
	if cache[M] == nil {
		cache[M] = make([]int, len(piles))
	}
	if cache[M][currentIndex] != 0 {
		return cache[M][currentIndex]
	}
	for i := currentIndex; i < currentIndex+M*2 && i < len(piles); i++ {
		cache[M][currentIndex] = max(
			cache[M][currentIndex],
			piles[currentIndex]-dp(cache, piles, max(M, i-currentIndex+1), i+1),
		)
	}
	return cache[M][currentIndex]
}

func stoneGameII(piles []int) int {
	sum := 0
	for i := len(piles) - 1; i >= 0; i-- {
		piles[i] += sum
		sum = piles[i]
	}
	cache := make([][]int, len(piles)+1)
	r := dp(cache, piles, 1, 0)
	return r
}

func TestGame(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		104,
		stoneGameII([]int{1, 2, 3, 4, 5, 100}),
	)
	assert.Equal(
		10,
		stoneGameII([]int{2, 7, 9, 4, 4}),
	)
}
