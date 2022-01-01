package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numSplits(s string) (times int) {
	left := make(map[byte]int)
	right := make(map[byte]int)
	for idx := range s {
		right[s[idx]] += 1
	}

	for idx := 0; idx < len(s) && len(left) <= len(right); idx++ {
		left[s[idx]] += 1
		right[s[idx]] -= 1
		if right[s[idx]] == 0 {
			delete(right, s[idx])
		}
		if len(left) == len(right) {
			times += 1
		}
	}
	return
}

func TestSplit(t *testing.T) {
	assert.Equal(t, 2, numSplits("aacaba"))
	assert.Equal(t, 1, numSplits("abcd"))
}
