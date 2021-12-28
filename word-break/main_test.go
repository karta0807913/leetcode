package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	candidates := make([]string, 1, len(wordDict))
	cache := make(map[string]bool)
	candidates[0] = s
	maxLength := 0
	for _, word := range wordDict {
		wordMap[word] = true
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	for len(candidates) != 0 {
		word := candidates[0]
		candidates = candidates[1:]
		for idx := 1; idx <= maxLength && idx <= len(word); idx++ {
			if wordMap[word[:idx]] {
				newWord := word[idx:]
				if newWord == "" {
					return true
				}
				if !cache[newWord] {
					cache[newWord] = true
					candidates = append(candidates, newWord)
				}

			}
		}
	}
	return false
}

func TestBreak(t *testing.T) {
	assert.True(t, wordBreak("leetcode", []string{"leet", "code"}))
	assert.True(t, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.False(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	assert.True(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat", "an"}))
}
