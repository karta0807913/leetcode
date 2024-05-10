package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	memory := make([][]int, len(arr1))
	for i := range memory {
		memory[i] = make([]int, len(arr2)+1)
		for j := range memory[i] {
			memory[i][j] = -1
		}
	}
	ans := recursive(memory, -1, arr1, arr2, 0, 0)
	if ans == len(arr2)+1 {
		return -1
	}
	return ans
}

func recursive(memory [][]int, prev int, arr1, arr2 []int, index1, index2 int) int {
}

func TestIncreasing(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		1,
		makeArrayIncreasing([]int{
			10, 100, 40, 70, 80, 90,
		}, []int{
			110, 120, 30, 130, 140,
		}),
	)
	assert.Equal(
		1,
		makeArrayIncreasing([]int{
			1, 5, 3, 6, 7,
		}, []int{
			1, 3, 2, 4,
		}),
	)
	assert.Equal(
		2,
		makeArrayIncreasing([]int{
			1, 5, 3, 6, 7,
		}, []int{
			4, 3, 1,
		}),
	)
	assert.Equal(
		5,
		makeArrayIncreasing([]int{
			0, 11, 6, 1, 4, 3,
		}, []int{
			5, 4, 11, 10, 1, 0,
		}),
	)
	assert.Equal(
		-1,
		makeArrayIncreasing([]int{
			1, 5, 3, 6, 7,
		}, []int{
			1, 6, 3, 3,
		}),
	)
	assert.Equal(
		-1,
		makeArrayIncreasing([]int{
			7, 6, 3, 0,
		}, []int{
			9,
		}),
	)
}
