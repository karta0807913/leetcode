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

func pushDominoes(dominoes string) string {
	dp := make([]int, len(dominoes))
	force := 0

	result := []byte(dominoes)
	for i := len(result) - 1; i >= 0; i-- {
		direct := result[i]
		switch direct {
		case 'L':
			force = len(result)
		case 'R':
			force = 0
		default:
			force = max(force-1, 0)
		}
		dp[i] -= force
	}

	for i, direct := range result {
		switch direct {
		case 'R':
			force = len(result)
		case 'L':
			force = 0
		default:
			force = max(force-1, 0)
		}
		dp[i] += force
		if dp[i] > 0 {
			result[i] = 'R'
		} else if dp[i] < 0 {
			result[i] = 'L'
		} else {
			result[i] = '.'
		}
	}
	return string(result)
}

func TestDominos(t *testing.T) {
	assert := require.New(t)
	assert.Equal("LL.RR.LLRRLL..", pushDominoes(".L.R...LR..L.."))
	assert.Equal("RR.LL", pushDominoes("R...L"))
	assert.Equal("RRRRRRRRRRRRRR", pushDominoes("R............."))
}
