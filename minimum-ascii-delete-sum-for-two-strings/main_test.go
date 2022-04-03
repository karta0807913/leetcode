package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumDeleteSum(s1 string, s2 string) int {

	dp := make([]int, len(s2)+1)
	next := make([]int, len(s2)+1)

	for i := range s2 {
		dp[i+1] = dp[i] + int(s2[i])
	}

	// fmt.Printf("dp: %v\n", dp)
	for i := range s1 {
		next[0] = dp[0] + int(s1[i])
		for idx := 1; idx <= len(s2); idx++ {
			if s1[i] == s2[idx-1] {
				next[idx] = dp[idx-1]
			} else {
				deleteAll := int(s1[i]) + int(s2[idx-1]) + dp[idx-1]
				deleteS1NewChar := int(s1[i]) + dp[idx]
				deleteS2NewChar := int(s2[idx-1]) + next[idx-1]
				minVal := deleteAll
				if deleteS1NewChar < minVal {
					minVal = deleteS1NewChar
				}
				if deleteS2NewChar < minVal {
					minVal = deleteS2NewChar
				}
				next[idx] = minVal
			}
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[len(s2)]
}

func TestMinSum(t *testing.T) {
	assert.Equal(t, 231, minimumDeleteSum("sea", "eat"))
	assert.Equal(t, 403, minimumDeleteSum("delete", "leet"))
	assert.Equal(t, 0, minimumDeleteSum("aaa", "aaa"))
	assert.Equal(t, 194, minimumDeleteSum("aaa", "a"))
	assert.Equal(t, 194, minimumDeleteSum("a", "aaa"))
}
