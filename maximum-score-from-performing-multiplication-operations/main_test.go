package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximumScore(nums []int, multipliers []int) int {
	dpLeft := make([]int, len(nums))
	dpRight := make([]int, len(nums))
	nextLeft := make([]int, len(nums))
	nextRight := make([]int, len(nums))

	//pre-process
	for i, mult := 0, multipliers[len(multipliers)-1]; i < len(nums); i++ {
		dpLeft[i], dpRight[i] = nums[i]*mult, nums[i]*mult
	}

	// fmt.Printf("dp: %v, %v\n", dpLeft, dpRight)
	// built from the bottom up
	// since len(nums) always bigger or equal than len(multipliers),
	//       we have to calculate the diff
	for length := len(nums) - len(multipliers) + 1; length < len(nums); length++ {
		// calculate the actual multipliers index
		//                  len(multipliers)-(the diff between nums and multipliers)-1
		mult := multipliers[len(multipliers)-(length-(len(nums)-len(multipliers)))-1]
		// if next position is left
		for leftIndex := 0; leftIndex < len(nums)-length; leftIndex++ {
			nextLeft[leftIndex] = max(
				// previous position is the rightest
				// c o o o p
				dpRight[leftIndex+length],
				//previous position is on the left
				// c p o o o
				dpLeft[leftIndex+1],
			)
			nextLeft[leftIndex] += nums[leftIndex] * mult
		}

		// if next position is right
		for rightIndex := len(nums) - 1; rightIndex >= length; rightIndex-- {
			nextRight[rightIndex] = max(
				// previous position is left
				// p o o o c
				dpLeft[rightIndex-length],
				// previous position is right
				// o o o p c
				dpRight[rightIndex-1],
			)
			nextRight[rightIndex] += nums[rightIndex] * mult
		}
		dpLeft, nextLeft = nextLeft, dpLeft
		dpRight, nextRight = nextRight, dpRight
		// fmt.Printf("dp: %v, %v, %v\n", dpLeft, dpRight, mult)
	}

	// max(last item on the left, last item on the right)
	return max(dpLeft[0], dpRight[len(nums)-1])
}

func TestScore(t *testing.T) {
	assert := require.New(t)

	assert.Equal(102, maximumScore([]int{
		-5, -3, -3, -2, 7, 1,
	}, []int{
		-10, -5, 3, 4, 6,
	}))
	assert.Equal(14, maximumScore([]int{
		1, 2, 3,
	}, []int{
		3, 2, 1,
	}))
}
