package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeExtraSpace(s []rune) []rune {
	spaceFlag := true
	startIdx, endIdx := 0, 0
	for endIdx < len(s) {
		if s[endIdx] != ' ' {
			spaceFlag = false
		}
		if s[endIdx] == ' ' && spaceFlag {
			endIdx += 1
			if endIdx == len(s) {
				break
			}
		} else {
			if s[endIdx] == ' ' {
				spaceFlag = true
			}
			s[endIdx], s[startIdx] = s[startIdx], s[endIdx]
			endIdx += 1
			startIdx += 1
		}
	}
	if s[startIdx-1] == ' ' {
		startIdx -= 1
	}
	return s[:startIdx]
}

func reverse(s []rune) []rune {
	for idx := 0; idx < len(s)/2; idx++ {
		s[idx], s[len(s)-idx-1] = s[len(s)-idx-1], s[idx]
	}
	return s
}

func reverseWords(s string) string {
	converted := []rune(s)
	reverse(converted)
	start := 0
	for idx, val := range converted {
		if val == ' ' {
			reverse(converted[start:idx])
			start = idx + 1
		}
	}
	reverse(converted[start:])
	return string(removeExtraSpace(converted))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []rune("word"), removeExtraSpace([]rune("   word")), nil)
	assert.Equal(t, []rune("word"), removeExtraSpace([]rune("   word   ")), nil)
	assert.Equal(t, []rune("hello word"), removeExtraSpace([]rune("  hello word   ")), nil)
	assert.Equal(t, []rune("hello word"), removeExtraSpace([]rune("  hello   word   ")), nil)
	assert.Equal(t, []rune("abcd"), reverse([]rune("dcba")), nil)
	assert.Equal(t, []rune("abcde"), reverse([]rune("edcba")), nil)
	assert.Equal(t, "abcd", reverseWords("abcd"), nil)
	assert.Equal(t, "ab cd", reverseWords("cd ab"), nil)
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"), nil)
	assert.Equal(t, "word hello", reverseWords("  hello word  "), nil)
	assert.Equal(t, "example good a", reverseWords("a good   example"), nil)
	assert.Equal(t, "Alice Loves Bob", reverseWords("  Bob    Loves  Alice   "), nil)
	assert.Equal(t, "bob like even not does Alice", reverseWords("Alice does not even like bob"), nil)
}
