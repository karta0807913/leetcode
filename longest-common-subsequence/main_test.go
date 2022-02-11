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

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	dp := make([]int, len(text2))
	next := make([]int, len(text2))
	for i1 := range text1 {
		for i2 := range text2 {
			if text1[i1] == text2[i2] {
				prev := 0
				if i2 != 0 {
					prev = dp[i2-1]
				}
				next[i2] = max(dp[i2], prev+1)
			} else {
				prev := 0
				if i2 != 0 {
					prev = next[i2-1]
				}
				next[i2] = max(dp[i2], prev)
			}
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v, text1[i1]: %c\n", dp, text1[i1])
	}
	return dp[len(dp)-1]
}

func TestLen(t *testing.T) {
	assert.Equal(t, 3, longestCommonSubsequence("abcde", "ace"))
	assert.Equal(t, 4, longestCommonSubsequence("abdde", "abcde"))
	assert.Equal(t, 4, longestCommonSubsequence("abcde", "abdde"))
	assert.Equal(t, 3, longestCommonSubsequence("abcde", "bdde"))
	assert.Equal(t, 1, longestCommonSubsequence("psnw", "vozsh"))
}
