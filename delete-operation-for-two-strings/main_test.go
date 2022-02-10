package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := range dp {
		dp[i] = i
	}
	next := make([]int, len(word2)+1)
	for i1 := range word1 {
		next[0] = dp[0] + 1
		for i2 := range word2 {
			if word1[i1] == word2[i2] {
				next[i2+1] = dp[i2]
			} else {
				next[i2+1] = min(next[i2], dp[i2+1]) + 1
			}
		}
		next, dp = dp, next
		// fmt.Printf("dp: %v, i: %c\n", dp, word1[i1])
	}
	return dp[len(dp)-1]
}

func TestMin(t *testing.T) {
	assert.Equal(t, 4, minDistance("leetcode", "etco"))
	assert.Equal(t, 4, minDistance("etco", "leetcode"))
	assert.Equal(t, 4, minDistance("abc", "aabbccc"))
}
