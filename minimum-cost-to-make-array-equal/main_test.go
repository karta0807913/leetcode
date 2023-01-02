package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func abs(a int) int64 {
	if a < 0 {
		return -int64(a)
	}
	return int64(a)
}

func minCost(nums []int, cost []int) int64 {
	weightSum := int64(0)
	sum := int64(0)
	for _, cost := range cost {
		weightSum += int64(cost)
	}
	for i, num := range nums {
		sum += int64(num) * int64(cost[i])
	}
	sum /= weightSum
	fmt.Printf("sum: %v\n", sum)
	// sum = 587000
	upperTarget := int(sum) + 1
	lowerTarget := int(sum)
	upperCost := int64(0)
	lowerCost := int64(0)
	for i, num := range nums {
		upperCost += abs(num-upperTarget) * int64(cost[i])
		lowerCost += abs(num-lowerTarget) * int64(cost[i])
	}
	if upperCost < lowerCost {
		return upperCost
	}
	return lowerCost
}

func TestCost(t *testing.T) {
	assert := require.New(t)
	assert.Equal(99+98, minCost([]int{
		1, 2, 100,
	}, []int{
		1, 1, 100,
	}))

	assert.Equal(1907611126748, minCost([]int{
		735103, 366367, 132236, 133334, 808160, 113001, 49051, 735598, 686615, 665317, 999793, 426087, 587000, 649989, 509946, 743518,
	}, []int{
		724182, 447415, 723725, 902336, 600863, 287644, 13836, 665183, 448859, 917248, 397790, 898215, 790754, 320604, 468575, 825614,
	}))

	assert.Equal(0, minCost([]int{
		2, 2, 2, 2, 2,
	}, []int{
		4, 2, 8, 1, 3,
	}))
	assert.Equal(8, minCost([]int{
		1, 3, 5, 2,
	}, []int{
		2, 3, 1, 14,
	}))

}
