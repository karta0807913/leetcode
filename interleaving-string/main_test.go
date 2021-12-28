package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 {
		return len(s1) == 0 && len(s2) == 0
	}
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	cache := make([]bool, len(s2)+1)
	cache[0] = true
	for s2Idx := range s2 {
		cache[s2Idx+1] = (s2[s2Idx] == s3[s2Idx] && cache[s2Idx])
	}
	for s1Idx := range s1 {
		cache[0] = cache[0] && (s1[s1Idx] == s3[s1Idx])
		for s2Idx := range s2 {
			cache[s2Idx+1] = (s2[s2Idx] == s3[s1Idx+s2Idx+1] && cache[s2Idx]) ||
				(s1[s1Idx] == s3[s1Idx+s2Idx+1] && cache[s2Idx+1])
		}
	}
	return cache[len(s2)]
}

// O(2^{m+n})
func isInterleaveRecursive(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s3) == 0 {
		return len(s1) == 0 && len(s2) == 0
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	if len(s1) != 0 && s1[0] == s3[0] {
		if isInterleaveRecursive(s1[1:], s2, s3[1:]) {
			return true
		}
	}

	if len(s2) != 0 && s2[0] == s3[0] {
		if isInterleaveRecursive(s1, s2[1:], s3[1:]) {
			return true
		}
	}

	return false
}

func TestInterleave(t *testing.T) {
	assert.True(t, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.False(t, isInterleave("aabcc", "abbca", "aadbbbaccc"))
	assert.True(t, isInterleave("", "", ""))
	assert.False(t, isInterleave("a", "", ""))
	assert.False(t, isInterleave("aaaa", "aa", "aaa"))
	assert.False(t, isInterleave("aa", "aaaa", "aaa"))
	assert.False(t, isInterleave("abababababababababababababababababababababababababababababababababababababababababababababababababbb", "babababababababababababababababababababababababababababababababababababababababababababababababaaaba", "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"))
}
