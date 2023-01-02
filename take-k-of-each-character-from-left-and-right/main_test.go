package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	cache := make(map[byte]int)
	for i := len(s) - 1; i >= 0; i-- {
		cache[s[i]]++
	}
	for _, b := range []byte{'a', 'b', 'c'} {
		if cache[b] < k {
			return -1
		}
	}

	cachedIndex := 0

	for ; cachedIndex < len(s) && cache[s[cachedIndex]] >= k+1; cachedIndex++ {
		cache[s[cachedIndex]]--
	}

	minLength := len(s) - cachedIndex

	for i := range s {
		cache[s[i]]++

		for ; cachedIndex < len(s) && cache[s[cachedIndex]] >= k+1; cachedIndex++ {
			cache[s[cachedIndex]]--
		}

		// fmt.Println(i, cachedIndex, cache)
		if n := len(s) - cachedIndex + i + 1; minLength > n {
			minLength = n
		}
	}
	return minLength
}

func TestChar(t *testing.T) {
	assert := require.New(t)
	assert.Equal(8, takeCharacters("aabaaaacaabc", 2))
	assert.Equal(3, takeCharacters("cbbac", 1))
	assert.Equal(0, takeCharacters("a", 0))
}
