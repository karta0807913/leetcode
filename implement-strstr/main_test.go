package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strStr(haystack string, needle string) int {
	idx := 0
	start_point := 0
	for cidx := 0; cidx < len(haystack); cidx++ {
		if len(needle) == idx {
			return cidx - len(needle)
		}
		if idx == 0 {
			start_point = cidx
		}
		if needle[idx] == haystack[cidx] {
			idx += 1
		} else {
			idx = 0
			cidx = start_point
		}
	}
	if len(needle) == idx {
		return len(haystack) - len(needle)
	}
	return -1
}

func TestStr(t *testing.T) {
	assert.Equal(t, 5, strStr("laaaall", "ll"), nil)
	assert.Equal(t, 2, strStr("hello", "ll"), nil)
	assert.Equal(t, -1, strStr("aaaaaaa", "bba"), nil)
	assert.Equal(t, 0, strStr("", ""), nil)
	assert.Equal(t, 0, strStr("laaaall", "la"), nil)
	assert.Equal(t, 4, strStr("mississippi", "issip"), nil)
}
