package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countSubstrings(s string, t string) int {
	if len(t) < len(s) {
		t, s = s, t
	}
	strMap := make(map[string]int)
	for i := len(s); i <= len(t); i++ {
		strMap[t[i-len(s):i]] += 1
	}

	result := 0
	for subsetLength := len(s); subsetLength > 0; subsetLength-- {

		for i := subsetLength; i <= len(s); i++ {
			for j := 0; j < subsetLength; j++ {
				for b := 'a'; b <= 'z'; b++ {
					result += strMap[s[i-subsetLength:i+j-subsetLength]+string(b)+s[i+j+1-subsetLength:i]]
				}
				result -= strMap[s[i-subsetLength:i]]
			}
		}

		newMap := make(map[string]int)
		for subset := range strMap {
			newMap[subset[:subsetLength-1]] += strMap[subset]
		}
		newMap[t[len(t)-subsetLength+1:]] += 1
		strMap = newMap
	}
	return result
}

func TestCount(t *testing.T) {
	assert.Equal(t, 4, countSubstrings("ab", "ac"))
	assert.Equal(t, 18, countSubstrings("abab", "acac"))
	assert.Equal(t, 4, countSubstrings("ab", "vb"))
	assert.Equal(t, 6, countSubstrings("aba", "baba"))
	assert.Equal(t, 3, countSubstrings("ab", "bb"))
	assert.Equal(t, 90, countSubstrings("computer", "computation"))
}
