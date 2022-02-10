package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	next := make([]int, len(word2)+1)
	for i := range dp {
		dp[i] = i
	}
	// fmt.Printf("dp: %v\n", dp)
	for i1 := range word1 {
		next[0] = dp[0] + 1
		for i2 := range word2 {
			if word1[i1] == word2[i2] {
				next[i2+1] = dp[i2]
			} else {
				insert := dp[i2+1] + 1
				replace := dp[i2] + 1
				del := next[i2] + 1
				min := insert
				if replace < min {
					min = replace
				}
				if del < min {
					min = del
				}
				next[i2+1] = min
			}
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v, word1[i1]: %c\n", dp, word1[i1])
	}
	return dp[len(dp)-1]
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
	assert.Equal(t, 1, minDistance("ho", "ro"))
}
