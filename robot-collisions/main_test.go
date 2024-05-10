package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	indexes := make([]int, len(positions))
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return positions[indexes[i]] < positions[indexes[j]]
	})
	var aliveRobots []int
	var robotLeftToRight []int
	for _, index := range indexes {
		if directions[index] == 'R' {
			robotLeftToRight = append(robotLeftToRight, index)
			continue
		}
		for len(robotLeftToRight) != 0 {
			last := robotLeftToRight[len(robotLeftToRight)-1]
			if healths[last] > healths[index] {
				healths[index] = 0
				healths[last]--
				break
			}
			if healths[last] == healths[index] {
				robotLeftToRight = robotLeftToRight[:len(robotLeftToRight)-1]
				healths[index] = 0
				break
			}
			healths[index]--
			robotLeftToRight = robotLeftToRight[:len(robotLeftToRight)-1]
		}
		if healths[index] != 0 {
			aliveRobots = append(aliveRobots, index)
		}
	}
	aliveRobots = append(aliveRobots, robotLeftToRight...)
	sort.Ints(aliveRobots)
	ans := make([]int, 0, len(aliveRobots))
	for _, index := range aliveRobots {
		ans = append(ans, healths[index])
	}
	return ans
}

func TestRobot(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{2, 17, 9, 15, 10},
		survivedRobotsHealths(
			[]int{5, 4, 3, 2, 1},
			[]int{2, 17, 9, 15, 10},
			"RRRRR",
		),
	)
}
