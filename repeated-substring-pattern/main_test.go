package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildable(s, substring string) bool {
	if len(s)%len(substring) != 0 {
		return false
	}
	for offset := 0; offset < len(s); offset += len(substring) {
		if substring != s[offset:len(substring)+offset] {
			return false
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	for idx := len(s) / 2; idx > 0; idx-- {
		if buildable(s[idx:], s[:idx]) {
			return true
		}
	}
	return false
}

func TestPattern(t *testing.T) {
	assert.True(t, repeatedSubstringPattern("abcabc"), nil)
	assert.False(t, repeatedSubstringPattern("aabaaba"), nil)
	assert.False(t, buildable("abba", "ab"), nil)
	assert.False(t, repeatedSubstringPattern("ababba"), nil)
	assert.True(t, repeatedSubstringPattern("abcabcabcabc"), nil)
	assert.False(t, repeatedSubstringPattern("aba"), nil)
}
