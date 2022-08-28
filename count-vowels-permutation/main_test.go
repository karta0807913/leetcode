package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func countVowelPermutation(n int) int {
	dp := [5]int{1, 1, 1, 1, 1}
	next := [5]int{0, 0, 0, 0, 0}
	for n > 1 {
		n--
		next[1] = (next[1] + dp[0]) % mod
		dp[0] = 0

		next[0] = (next[0] + dp[1]) % mod
		next[2] = (next[2] + dp[1]) % mod
		dp[1] = 0

		next[0] = (next[0] + dp[2]) % mod
		next[1] = (next[1] + dp[2]) % mod
		next[3] = (next[3] + dp[2]) % mod
		next[4] = (next[4] + dp[2]) % mod
		dp[2] = 0

		next[2] = (next[2] + dp[3]) % mod
		next[4] = (next[4] + dp[3]) % mod
		dp[3] = 0

		next[0] = (next[0] + dp[4]) % mod
		dp[4] = 0

		dp, next = next, dp
	}
	result := 0
	for _, val := range dp {
		result = (result + val) % mod
	}
	return result
}

func TestVal(t *testing.T) {
	assert := require.New(t)
	assert.Equal(10, countVowelPermutation(2))
	assert.Equal(19, countVowelPermutation(3))
}
