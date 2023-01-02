package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type Point struct {
	Step int
	K    int
}

type Pair struct {
	X int
	Y int
}

type BlockStatus int

const (
	Blocked = iota
	Space
	OOB
)

func checkAndBlock(grid [][]int, status [][]Point, current, next Pair, k int) BlockStatus {
	if next.Y < 0 || len(status) <= next.Y {
		return OOB
	}
	if next.X < 0 || len(status[next.Y]) <= next.X {
		return OOB
	}
	currentPoint := &status[current.Y][current.X]
	currentPoint.K = k
	nextPoint := &status[next.Y][next.X]
	if nextPoint.Step == 0 || nextPoint.Step == -1 {
		nextPoint.Step = currentPoint.Step + 1
		nextPoint.K = currentPoint.K
		if grid[next.Y][next.X] == 0 {
			return Space
		}
		return Blocked
	}

	if nextPoint.Step <= currentPoint.Step+1 {
		return OOB
	}

	nextPoint.Step = currentPoint.Step + 1
	if nextPoint.K == k {
		return OOB
	}
	nextPoint.K = k
	if grid[next.Y][next.X] == 0 {
		return Space
	}
	return Blocked
}

func shortestPath(grid [][]int, k int) int {
	status := make([][]Point, len(grid))
	for i := range status {
		status[i] = make([]Point, len(grid[i]))
	}
	status[len(grid)-1][len(status[0])-1].Step = -1
	status[0][0].Step = 0
	startPoints := []Pair{{0, 0}}
	var blockedPoint, nextPoints []Pair
	for _k := 0; _k <= k; _k++ {
		for len(startPoints) != 0 {
			for len(startPoints) != 0 {
				current := startPoints[len(startPoints)-1]
				startPoints = startPoints[:len(startPoints)-1]
				for _, offset := range []Pair{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
					next := Pair{X: current.X + offset.X, Y: current.Y + offset.Y}
					switch checkAndBlock(grid, status, current, next, _k) {
					case Blocked:
						blockedPoint = append(blockedPoint, next)
					case Space:
						nextPoints = append(nextPoints, next)
					case OOB:
					}
				}
			}
			sort.Slice(nextPoints, func(i, j int) bool {
				return status[nextPoints[i].Y][nextPoints[i].X].Step <
					status[nextPoints[j].Y][nextPoints[j].X].Step
			})
			startPoints, nextPoints = nextPoints, startPoints[:0]
		}
		startPoints, blockedPoint = blockedPoint, startPoints
	}
	return status[len(grid)-1][len(status[0])-1].Step
}

func TestPath(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		6,
		shortestPath([][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		}, 1),
	)

	assert.Equal(
		-1,
		shortestPath([][]int{
			{0, 1, 1},
			{1, 1, 1},
			{1, 0, 0},
		}, 1),
	)

	assert.Equal(
		0,
		shortestPath([][]int{
			{0},
		}, 1),
	)
}
