package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	return (pow(a*a%mod, b/2) * pow(a, b%2) % mod)
}

func getInv(a int) int {
	return pow(a, mod-2)
}

func modDevided(left, right int) int {
	return (left * getInv(right) % mod)
}

func recursive(nums []int, mutiplies []int) int {
	if len(nums) < 3 {
		return 1
	}
	var left, right []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	// fmt.Println(left, right)
	return (modDevided(
		modDevided(
			mutiplies[len(left)+len(right)],
			mutiplies[len(left)],
		),
		mutiplies[len(right)],
	) * (recursive(
		left, mutiplies,
	) * recursive(
		right, mutiplies,
	) % mod) % mod)
}

func numOfWays(nums []int) int {
	mutiplies := make([]int, len(nums)+1)
	mutiplies[0] = 1
	for i := 1; i < len(mutiplies); i++ {
		mutiplies[i] = (mutiplies[i-1] * i) % mod
	}
	// fmt.Println(mutiplies)
	return recursive(nums, mutiplies) - 1
}

func TestWays(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		19,
		numOfWays([]int{3, 1, 2, 5, 4, 6}),
	)
	assert.Equal(
		0,
		numOfWays([]int{1, 2, 3}),
	)
	assert.Equal(
		79,
		numOfWays([]int{4, 2, 1, 3, 6, 5, 7}),
	)
}
