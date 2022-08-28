package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func minDeletionSize(strs []string) int {
	deleted := 0
	smaller := make([]bool, len(strs)-1)
	next := make([]bool, 0, len(strs)-1)
	for c := 0; c < len(strs[0]); c++ {
		for i := 0; i < len(strs)-1; i++ {
			if strs[i][c] > strs[i+1][c] && !smaller[i] {
				deleted += 1
				break
			}
			next = append(next, smaller[i] || strs[i][c] < strs[i+1][c])
		}
		if len(next) == len(smaller) {
			next, smaller = smaller, next
		}
		next = next[:0]
	}
	return deleted
}

func TestDeletion(t *testing.T) {
	assert := require.New(t)
	assert.Equal(0, minDeletionSize([]string{
		"xc", "yb", "za",
	}))
	assert.Equal(2, minDeletionSize([]string{
		"vdy", "vei", "zvc", "zld",
	}))
	assert.Equal(1, minDeletionSize([]string{
		"ca", "bb", "ac",
	}))
	assert.Equal(3, minDeletionSize([]string{
		"zyx", "wvu", "tsr",
	}))
	assert.Equal(0, minDeletionSize([]string{
		"abcdef", "uvwxyz",
	}))
}
