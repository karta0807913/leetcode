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

func stoneGameIII(stoneValue []int) string {
	sum := 0
	for i := len(stoneValue) - 1; i >= 0; i-- {
		stoneValue[i] += sum
		sum = stoneValue[i]
	}
	stoneValue = append(stoneValue, 0)
	for i := len(stoneValue) - 2; i >= 0; i-- {
		total := stoneValue[i]
		stoneValue[i] = total - stoneValue[i+1]
		if i+2 < len(stoneValue) {
			stoneValue[i] = max(stoneValue[i], total-stoneValue[i+2])
		}
		if i+3 < len(stoneValue) {
			stoneValue[i] = max(stoneValue[i], total-stoneValue[i+3])
		}
	}
	if sum < stoneValue[0]*2 {
		return "Alice"
	}
	if sum > stoneValue[0]*2 {
		return "Bob"
	}
	return "Tie"
}

func TestSt(t *testing.T) {
	assert := require.New(t)
	assert.Equal("Bob", stoneGameIII([]int{1, 2, 3, 7}))
	assert.Equal("Tie", stoneGameIII([]int{-1, -2, -3}))
}
