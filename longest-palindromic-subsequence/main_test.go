package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseq(s string) int {
	dp := make([]int, len(s))
	next := make([]int, len(s))
	for start := len(s) - 1; start >= 0; start-- {
		next[start] = 1
		for end := start + 1; end < len(s); end++ {
			if s[start] == s[end] {
				next[end] = dp[end-1] + 2
			} else {
				next[end] = max(dp[end], next[end-1])
			}
		}
		dp, next = next, dp
	}
	return dp[len(dp)-1]
}

func TestSubseq(t *testing.T) {
	assert.Equal(t, 1, longestPalindromeSubseq("a"))
	assert.Equal(t, 2, longestPalindromeSubseq("aa"))
	assert.Equal(t, 2, longestPalindromeSubseq("aab"))
	assert.Equal(t, 4, longestPalindromeSubseq("baab"))
	assert.Equal(t, 3, longestPalindromeSubseq("ababb"))
	assert.Equal(t, 5, longestPalindromeSubseq("ababa"))
	assert.Equal(t, 5, longestPalindromeSubseq("abcbab"))
	assert.Equal(t, 2, longestPalindromeSubseq("cbbd"))
}
