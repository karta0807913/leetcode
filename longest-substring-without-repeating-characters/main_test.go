package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func lengthOfLongestSubstring(str string) int {
	existsMap := make(map[byte]int, 52)
	maxLength := 0
	length := 0
	duplicated := false
	for i := 0; i < len(str); {
		if duplicated {
			existsMap[str[i-length]] -= 1
			if existsMap[str[i-length]] == 1 {
				duplicated = false
			}
			length--
		} else {
			existsMap[str[i]] += 1
			duplicated = existsMap[str[i]] != 1
			length++
			if !duplicated && maxLength < length {
				maxLength = length
			}
			i++
		}
	}
	return maxLength
}

func TestLength(t *testing.T) {
	assert := require.New(t)
	assert.Equal(3, lengthOfLongestSubstring("pwwkew"))
}
