package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cache []int = []int{1, 0, 0, 0, 0, 0, 0, 0}

func cal(counts []int) int {
	dev := 1
	length := 0
	for _, val := range counts {
		if val != 0 {
			length += val
			dev *= cache[val]
		}
	}
	if length == 0 {
		return 0
	}
	return cache[length] / dev
}

func recursive(counts []int, idx int) int {
	target := 0
	for idx = idx + 1; idx < len(counts); idx++ {
		if counts[idx] != 0 {
			target = idx
			break
		}
	}
	fmt.Println(counts, target)
	result := cal(counts)
	fmt.Println(result)
	if idx == len(counts) {
		return result
	}
	tmp := counts[target]
	for counts[target] > 0 {
		counts[target]--
		result += recursive(counts, target)
	}
	counts[target] = tmp
	return result
}

func numTilePossibilities(tiles string) int {
	for idx := 1; idx < len(cache); idx++ {
		cache[idx] = cache[idx-1] * idx
	}
	counts := make([]int, 26)
	for _, val := range tiles {
		counts[val-'A']++
	}
	return recursive(counts, -1)
}

func TestRecursive(t *testing.T) {
	for idx := 1; idx < len(cache); idx++ {
		cache[idx] = cache[idx-1] * idx
	}
	assert.Equal(t, 3, cal([]int{2, 1}), nil)
	// assert.Equal(t, 8, numTilePossibilities("AAB"), nil)
	assert.Equal(t, 18, numTilePossibilities("AABB"), nil)
	// assert.Equal(t, 188, numTilePossibilities("AAABBC"), nil)
}
