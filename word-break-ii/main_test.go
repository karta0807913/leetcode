package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Answer struct {
	prefix  []string
	remains string
}

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool, len(wordDict))
	candidates := make([]Answer, 1, len(wordDict))
	answers := make([]string, 0)
	candidates[0] = Answer{remains: s}
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
		for idx := 1; idx <= maxLength && idx <= len(word.remains); idx++ {
			if wordMap[word.remains[:idx]] {
				newPrefix := append(word.prefix, word.remains[:idx])
				newRemains := word.remains[idx:]
				if newRemains == "" {
					answers = append(answers, strings.Join(newPrefix, " "))
					break
				} else {
					candidates = append(candidates, Answer{
						// create a new array. if we don't do this, the memory space will messed up,
						// and the second and third test will fail.
						prefix:  append(make([]string, 0, len(newPrefix)), newPrefix...),
						remains: newRemains,
					})
				}
			}
		}
		// fmt.Printf("candidates: %v, answers: %v\n", candidates, answers)
	}

	return answers
}

func TestBreak(t *testing.T) {
	assert.ElementsMatch(t, []string{"cats and dog", "cat sand dog"}, wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	assert.ElementsMatch(t, []string{"cats an dog"}, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat", "an"}))

	assert.ElementsMatch(t, []string{
		"a a a a a", "aa a a a", "a aa a a", "a a aa a", "aa aa a", "aaaa a", "a a a aa", "aa a aa", "a aa aa", "a aaaa",
	}, wordBreak("aaaaa", []string{"aaaa", "aa", "a"}))
	assert.ElementsMatch(t, []string{
		"a a a a a", "aa a a a", "a aa a a", "a a aa a", "aa aa a", "a a a aa", "aa a aa", "a aa aa",
	}, wordBreak("aaaaa", []string{"aa", "a"}))
}
