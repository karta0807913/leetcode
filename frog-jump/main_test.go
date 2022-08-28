package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func key(i, k int) string {
	return fmt.Sprintf("%d,%d", i, k)
}

func recursive(stones map[int]bool, i, k int, cache map[string]bool, maxVal int) bool {
	if !stones[i] {
		return false
	}
	if i == maxVal {
		return true
	}
	key := key(i, k)
	if cache, ok := cache[key]; ok {
		return cache
	}
	// fmt.Printf("key: %v\n", key)
	cache[key] = recursive(stones, i+k+1, k+1, cache, maxVal) ||
		recursive(stones, i+k, k, cache, maxVal)
	if k-1 != 0 {
		cache[key] = cache[key] || recursive(stones, i+k-1, k-1, cache, maxVal)
	}
	// fmt.Printf("cache[key]: %v\n", cache[key])
	return cache[key]
}

func canCross(stones []int) bool {
	s := make(map[int]bool)
	for _, i := range stones {
		s[i] = true
	}
	return recursive(s, 1, 1, make(map[string]bool), stones[len(stones)-1])
}

func TestCross(t *testing.T) {
	assert := require.New(t)
	assert.False(canCross([]int{0, 2}))
	assert.True(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	assert.False(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}
