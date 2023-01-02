package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func numWays(words []string, target string) int {
	dp := make([]int, len(words[0])+1)
	next := make([]int, len(words[0])+1)
	wordsMap := make([][26]int, len(words[0]))
	for i := range wordsMap {
		wordsMap[i] = [26]int{}
		dp[i] = 1
		for _, word := range words {
			wordsMap[i][word[i]-'a'] += 1
		}
	}
	dp[len(dp)-1] = 1
	for idx := range target {
		for i := range words[0] {
			next[i+1] = next[i]
			next[i+1] += dp[i] * wordsMap[i][target[idx]-'a']
			next[i+1] %= mod
		}
		dp, next = next, dp
		next[0] = 0
	}
	return dp[len(dp)-1]
}

func TestWays(t *testing.T) {
	assert := require.New(t)
	assert.Equal(4, numWays([]string{
		"abba",
		"baab",
	}, "bab"))
	assert.Equal(3, numWays([]string{
		"ab",
		"ab",
		"ab",
	}, "a"))
	assert.Equal(9, numWays([]string{
		"ab",
		"ab",
		"ab",
	}, "ab"))
}
