package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod int = 1e9 + 7

func numberOfArrays(s string, k int) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		var n int
		for j := i; j < len(s); j++ {
			n *= 10
			n += int(s[j] - '0')
			if n > k {
				break
			}
			dp[i] += dp[j+1]
			dp[i] %= mod
		}
	}
	return dp[0]
}

func TestNumberOf(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		34,
		numberOfArrays("1234567890", 90),
	)
}
