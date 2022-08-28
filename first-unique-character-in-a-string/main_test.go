package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func firstUniqChar(s string) int {
	repeatMap := make([]int, 26)
	for _, c := range s {
		repeatMap[int(c-'a')] += 1
	}
	for i, c := range s {
		if repeatMap[int(c-'a')] == 1 {
			return i
		}
	}
	return -1
}

func TestFirstUniqChar(t *testing.T) {
	assert := require.New(t)
	assert.Equal(0, firstUniqChar("leetcode"))
}
