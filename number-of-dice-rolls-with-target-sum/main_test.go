package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e+9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sumUp(arr []int, proper int) {
	for i := range arr {
		arr[i] += proper
		arr[i] %= mod
	}
}

func numRollsToTarget(n int, k int, target int) int {
	diceMap := make([]int, target+1)
	nextMap := make([]int, target+1)

	diceMap[0] = 1
	lower := 0
	for ; n != 0; n-- {
		for i := lower; i != target+1 && diceMap[i] != 0; i++ {
			sumUp(nextMap[i+1:min(i+k+1, len(nextMap))], diceMap[i])
			diceMap[i] = 0
		}
		lower++
		diceMap, nextMap = nextMap, diceMap
	}
	return diceMap[target]
}

func TestRolls(t *testing.T) {
	assert := require.New(t)
	assert.Equal(2, numRollsToTarget(2, 6, 3))
}
