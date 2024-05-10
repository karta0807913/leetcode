package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partitionLabels(s string) []int {
	result := []int{}
	hashMap := [26]int{}
	for i := range s {
		hashMap[s[i]-'a'] = i + 1
	}
	idx := 0
	for idx < len(s) {
		start := idx
		end := idx + 1
		for ; start < end; start++ {
			end = max(hashMap[s[start]-'a'], end)
		}
		result = append(result, end-idx)
		idx = end
	}
	return result
}

func TestLabels(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{9, 7, 8},
		partitionLabels("ababcbacadefegdehijhklij"),
	)
	assert.Equal(
		[]int{10},
		partitionLabels("eccbbbbdec"),
	)
}
