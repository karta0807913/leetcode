package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([]int, len(s)+1)
	next := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = 1
	}
	for tIdx := range t {
		ways := 0
		for sIdx := range s {
			if s[sIdx] == t[tIdx] {
				ways += dp[sIdx]
			}
			next[sIdx+1] = ways
		}
		if ways == 0 {
			return 0
		}
		dp, next = next, dp
		dp[0] = 0
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[len(dp)-1]
}

func TestNum(t *testing.T) {
	assert.Equal(t, 3, numDistinct("bccaba", "ba"))
	assert.Equal(t, 1, numDistinct("eee", "eee"))
}
