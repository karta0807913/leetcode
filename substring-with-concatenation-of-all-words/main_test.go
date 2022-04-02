package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func cal(s string, wordsMap map[string]int, l int) bool {
	ans := 0
	tmp := make(map[string]int)
	for i := l; i <= len(s); i += l {
		t := s[i-l : i]
		p, ok := tmp[t]
		if ok {
			if p == 0 {
				return false
			}
			ans += 1
			tmp[t] -= 1
		} else {
			c, k := wordsMap[t]
			if !k {
				return false
			}
			ans += 1
			tmp[t] = c - 1
		}
	}
	return true
}

func findSubstring(s string, words []string) (ans []int) {
	wordsMap := make(map[string]int)
	for _, w := range words {
		wordsMap[w] += 1
	}
	totalLen := len(words) * len(words[0])
	for idx := totalLen; idx <= len(s); idx += 1 {
		r := s[idx-totalLen : idx]
		if cal(r, wordsMap, len(words[0])) {
			ans = append(ans, idx-totalLen)
		}
	}
	return ans
}

func TestCal(t *testing.T) {
	wordsMap := make(map[string]int)
	wordsMap["good"] += 2
	wordsMap["word"] += 1
	wordsMap["best"] += 1
	assert.True(t, cal("goodgoodbestword", wordsMap, 4))

	assert.Equal(t, []int{8},
		findSubstring("wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "good"}))
}
