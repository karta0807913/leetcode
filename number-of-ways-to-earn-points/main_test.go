package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func waysToReachTarget(target int, types [][]int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		for i := len(dp) - 1; i >= t[1]; i-- {
			if prevCount := dp[i-t[1]]; prevCount != 0 {
				for next, count := i, t[0]; next < len(dp) && count != 0; next, count = next+t[1], count-1 {
					dp[next] += prevCount
					dp[next] %= mod
				}
			}
			// fmt.Printf("\tdp: %v, t: %v\n", dp, t)
		}
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[len(dp)-1]
}

func TestTarget(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		7,
		waysToReachTarget(
			6,
			[][]int{
				{6, 1}, {3, 2}, {2, 3},
			},
		),
	)

	assert.Equal(
		1,
		waysToReachTarget(
			18,
			[][]int{
				{6, 1}, {3, 2}, {2, 3},
			},
		),
	)
}
