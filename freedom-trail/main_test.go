package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type Test struct {
	AA int `b:"a"`
	B  struct {
		C int
	}
	D struct {
	}
}

func findRotateSteps(ring string, key string) int {
	current, next := make([]int, len(ring)), make([]int, len(ring))
	var graph [26][]int
	for idx, c := range ring {
		key := int(c - 'a')
		graph[key] = append(graph[key], idx)
	}
	previousIdx := []int{0}
	for _, c := range key {
		currentKey := int(c - 'a')
		for _, nextIdx := range graph[currentKey] {
			next[nextIdx] = math.MaxInt
			for _, parent := range previousIdx {
				next[nextIdx] = min(
					next[nextIdx],
					min(
						(parent-nextIdx+len(ring))%len(ring)+1,
						(nextIdx-parent+len(ring))%len(ring)+1,
					)+current[parent],
				)
			}
		}
		current, next = next, current
		previousIdx = graph[currentKey]
	}
	ans := math.MaxInt
	for _, idx := range previousIdx {
		ans = min(ans, current[idx])
	}
	// fmt.Printf("current: %v\n", current)
	return ans
}

func TestFindRotateSteps(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		4,
		findRotateSteps("godding", "gd"),
	)
	assert.Equal(
		493,
		findRotateSteps("kczyfdzrbzctgqrwrlcohizjb", "tqthzroylybfztczgcfkcbzyibzkzzzdbrcrbfcjrqjrlwyozwzikghizcowgbcrqckzrtjhrlzcrcrdbdcjzbfdhoirzqwrclgz"),
	)
	assert.Equal(
		5,
		findRotateSteps("aaaaa", "aaaaa"),
	)
	assert.Equal(
		13,
		findRotateSteps("godding", "godding"),
	)
}
