package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func DistanceToZero(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func binarySearch(points [][]int, distance int) (splitPoint int) {
	leftPtr, rightPtr := 0, len(points)-1
	for leftPtr <= rightPtr {
		if DistanceToZero(points[leftPtr]) < distance {
			leftPtr++
		} else {
			points[leftPtr], points[rightPtr] =
				points[rightPtr], points[leftPtr]
			rightPtr--
		}
	}
	fmt.Printf("points: %v, distance: %v\n", points, distance)
	return leftPtr
}

func kClosest(points [][]int, k int) [][]int {
	min, max := math.MaxInt, math.MinInt
	for _, point := range points {
		distance := DistanceToZero(point)
		if min > distance {
			min = distance
		}
		if max < distance {
			max = distance
		}
	}
	region := points
	remains := k
	for min <= max {
		mid := min + (max-min)/2
		m := binarySearch(region, mid)
		if m < remains {
			remains -= m
			region = region[m:]
			min = mid + 1
		} else if m > remains {
			max = mid - 1
		} else {
			break
		}
	}
	fmt.Printf("points: %v\n", points)
	return points[:k]
}

func TestCloset(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch([][]int{
		{64, -18}, {-34, -224},
	}, kClosest([][]int{
		{64, -18}, {-34, -224}, {-134, 193},
	}, 2))
	assert.ElementsMatch([][]int{
		{-2, 2},
	}, kClosest([][]int{
		{1, 3}, {-2, 2},
	}, 1))
	assert.ElementsMatch([][]int{
		{3, 3}, {-2, 4},
	}, kClosest([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
	assert.ElementsMatch([][]int{
		{3, 3}, {-2, 4},
	}, kClosest([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
}
